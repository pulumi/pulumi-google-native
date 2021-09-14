package gen

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/blang/semver"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/zclconf/go-cty/cty"

	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"

	"github.com/aymerick/raymond"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"gopkg.in/yaml.v2"

	dstar "github.com/bmatcuk/doublestar/v4"
)

type Sample struct {
	Description string
	Body        map[string]interface{}
}

type SampleMetadata struct {
	Description  string        `yaml:"description,omitempty"`
	Name         string        `yaml:"name,omitempty"`
	Resource     string        `yaml:"resource,omitempty"`
	Updates      []interface{} `yaml:"updates,omitempty"`
	Variables    []Variables   `yaml:"variables,omitempty"`
	Versions     []string      `yaml:"versions,omitempty"`
	Dependencies []string      `yaml:"dependencies,omitempty"`
}

type Variables struct {
	Name string `yaml:"name,omitempty"`
	Type string `yaml:"type,omitempty"`
}

type tokenGrouping struct {
	beta map[string]string
	ga   map[string]string
}

// ExtractSamples returns the samples for each resource token.
func ExtractSamples(baseFs fs.FS, pkgSpec *schema.PackageSpec, meta *resources.CloudAPIMetadata) (map[string][]Sample, error) {
	const basePattern = "services/google/**/samples/*.yaml"

	serviceGrouping, err := serviceGrouping(pkgSpec)
	if err != nil {
		return nil, err
	}

	samples := map[string][]Sample{}
	if err := dstar.GlobWalk(baseFs, basePattern, func(path string, d fs.DirEntry) error {
		s, err := parseSample(baseFs, path, serviceGrouping, meta)
		if err != nil {
			fmt.Printf("Failure processing sample: %q. Error: %v\n", path, err)
			return nil
		}

		if len(s) > 0 {
			fmt.Printf("Parsed %q:\n%#v\n", path, s)
		}

		for k, v := range s {
			if existing, has := samples[k]; has {
				existing = append(existing, v...)
			} else {
				samples[k] = v
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return samples, nil
}

func serviceGrouping(pkgSpec *schema.PackageSpec) (map[string]tokenGrouping, error) {
	serviceGrouping := map[string]tokenGrouping{}
	for t := range pkgSpec.Resources {
		tok, err := tokens.ParseTypeToken(t)
		if err != nil {
			return nil, err
		}
		pkg := tok.Module().Name()
		service, ver := filepath.Split(pkg.String())
		service, ver = filepath.Clean(service), filepath.Clean(ver)
		tokGrp, has := serviceGrouping[service]
		if !has {
			tokGrp = tokenGrouping{
				beta: map[string]string{},
				ga:   map[string]string{},
			}
		}
		if strings.Contains(ver, "beta") {
			tokGrp.beta[ver] = t
		} else {
			tokGrp.ga[ver] = t
		}
		serviceGrouping[service] = tokGrp
	}
	return serviceGrouping, nil
}

func parseSample(fs fs.FS, path string, grouping map[string]tokenGrouping, meta *resources.CloudAPIMetadata) (map[string][]Sample, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	metadata := SampleMetadata{}
	if err := yaml.NewDecoder(f).Decode(&metadata); err != nil {
		return nil, fmt.Errorf("parsing failure for %q: %w", path, err)
	}
	if metadata.Resource == "" {
		return nil, fmt.Errorf("no resource specified in %q", path)
	}
	if len(metadata.Dependencies) > 0 {
		// TODO: Handle dependencies.
		fmt.Printf("Skipping %q. It uses dependencies which are not yet supported.\n", path)
		return nil, nil
	}
	metadata.Resource = strings.ReplaceAll(metadata.Resource, "samples/", "")
	baseDir := filepath.Dir(path)
	sampleFile := filepath.Join(baseDir, metadata.Resource)
	s, err := fs.Open(sampleFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = s.Close()
	}()

	b, err := io.ReadAll(s)
	if err != nil {
		return nil, err
	}

	// Perform template replacement
	variables := map[string]string{}
	for _, v := range metadata.Variables {
		variables[v.Name] = v.Type
	}
	t, err := raymond.Parse(string(b))
	if err != nil {
		return nil, err
	}

	buf, err := t.Exec(variables)
	if err != nil {
		return nil, fmt.Errorf("failed to update template with variables: %#v: %w", variables, err)
	}

	raw := map[string]interface{}{}
	if err := json.NewDecoder(strings.NewReader(buf)).Decode(&raw); err != nil {
		return nil, err
	}

	var hasRefSamples func(interface{}) bool
	hasRefSamples = func(in interface{}) bool {
		switch reflect.TypeOf(in).Kind() {
		case reflect.Map:
			asMap := in.(map[string]interface{})
			for _, v := range asMap {
				return hasRefSamples(v)
			}
		case reflect.Slice, reflect.Array:
			asSlice := in.([]interface{})
			for _, i := range asSlice {
				return hasRefSamples(i)
			}
		case reflect.String:
			if strings.HasPrefix(in.(string), "{{ref:") {
				return true
			}
		}
		return false
	}

	if hasRefSamples(raw) {
		fmt.Printf("%q has reference field in variables, ignoring", path)
		return nil, nil
	}

	// Extract token(s)
	baseName := filepath.Base(metadata.Resource)
	ext := filepath.Ext(baseName)
	withoutExt := baseName
	if ext != "" {
		withoutExt = baseName[:len(baseName)-len(ext)]
	}
	splitResource := strings.Split(withoutExt, ".")
	if len(splitResource) < 2 {
		return nil, fmt.Errorf("malformed resource item: %q for %q", metadata.Resource, path)
	}

	sp := strings.Split(path, string(os.PathSeparator))
	serviceName := sp[len(sp)-3]
	resourceName := camel(splitResource[1])

	grp, has := grouping[serviceName]
	if !has {
		return nil, fmt.Errorf("missing grouping information for service: %q", serviceName)
	}
	findTokens := func(tokGrp tokenGrouping, serviceName, resourceName, version string) []string {
		var toks []string
		switch version {
		case "beta":
			for _, t := range grp.beta {
				typTok, err := tokens.ParseTypeToken(t)
				contract.AssertNoError(err)
				if strings.EqualFold(typTok.Name().String(), resourceName) {
					toks = append(toks, t)
				}
			}
		case "ga":
			for _, t := range grp.ga {
				typTok, err := tokens.ParseTypeToken(t)
				contract.AssertNoError(err)
				if strings.EqualFold(typTok.Name().String(), resourceName) {
					toks = append(toks, t)
				}
			}
		}
		return toks
	}

	samples := map[string][]Sample{}

	for _, v := range metadata.Versions {
		toks := findTokens(grp, serviceName, resourceName, v)
		for _, tok := range toks {
			sample := Sample{
				Description: metadata.Description,
			}
			res := meta.Resources[tok]
			sample.Body = transformBody(res, raw)
			samples[tok] = []Sample{sample}
		}
	}

	return samples, nil
}

func transformBody(res resources.CloudAPIResource, sample map[string]interface{}) map[string]interface{} {
	inputsMap := map[string]interface{}{}
	for name, prop := range res.CreateProperties {
		parent := sample
		if prop.Container != "" {
			container, ok := sample[prop.Container].(map[string]interface{})
			if !ok {
				continue
			}
			parent = container
		}

		if value, has := parent[name]; has {
			sdkName := name
			if prop.SdkName != "" {
				sdkName = prop.SdkName
			}
			inputsMap[sdkName] = value
		}
	}
	return inputsMap
}

type description string
type programText string
type language string

type languageToExampleProgram map[language]programText
type exampleRenderData struct {
	ExampleDescription       description
	LanguageToExampleProgram languageToExampleProgram
}
type resourceExamplesRenderData struct {
	Data []exampleRenderData
}

func SampleCode(pkgSpec *schema.PackageSpec, resExamples map[string][]Sample, languages []string) error {
	sortedKeys := codegen.SortedKeys(pkgSpec.Resources) // To generate in deterministic order

	// cache to speed up code generation
	hcl2Cache := hcl2.Cache(hcl2.NewPackageCache())
	pkg, err := schema.ImportSpec(*pkgSpec, nil)
	if err != nil {
		return err
	}
	loaderOption := hcl2.Loader(InMemoryPackageLoader(map[string]*schema.Package{
		"google-native": pkg,
	}))

	for _, pulumiToken := range sortedKeys {
		if ShouldExclude(pulumiToken) {
			fmt.Printf("Skipping %q since it matches exclusion pattern\n", pulumiToken)
			continue
		}

		split := strings.Split(pulumiToken, ":")
		if len(split) == 0 {
			return fmt.Errorf("invalid resourcename: %s", pulumiToken)
		}
		resourceName := ToUpperCamel(split[len(split)-1])

		examplesRenderData := resourceExamplesRenderData{}
		if resourceExamples, ok := resExamples[pulumiToken]; ok {
			for _, example := range resourceExamples {
				var items []model.BodyItem

				exampleJSON := example.Body
				keys := codegen.SortedKeys(exampleJSON)
				for _, k := range keys {
					v := exampleJSON[k]
					val, err := RenderValue(v)
					if err != nil {
						return err
					}
					items = append(items, &model.Attribute{
						Name:  k,
						Value: val,
					})
				}

				block := model.Block{
					Type:   "resource",
					Body:   &model.Body{Items: items},
					Labels: []string{resourceName, pulumiToken},
				}
				body := &model.Body{Items: []model.BodyItem{&block}}
				FormatBody(body)
				languageExample, err := generateExamplePrograms(pulumiToken, example, body, languages, hcl2Cache, loaderOption)
				if err != nil {
					return err
				}

				examplesRenderData.Data = append(examplesRenderData.Data,
					exampleRenderData{
						ExampleDescription:       description(example.Description),
						LanguageToExampleProgram: languageExample,
					})
			}
			if len(examplesRenderData.Data) > 0 {
				err := renderExampleToSchema(pkgSpec, pulumiToken, &examplesRenderData)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func renderExampleToSchema(pkgSpec *schema.PackageSpec, resourceName string,
	examplesRenderData *resourceExamplesRenderData) error {
	const tmpl = `

{{"{{% examples %}}"}}
## Example Usage
{{- range .Data }}
{{ "{{% example %}}" }}
### {{ .ExampleDescription }}

{{- range $lang, $example := .LanguageToExampleProgram }}
{{ beginLanguage $lang }}
{{ $example }}
{{ endLanguage }}
{{ end }}
{{"{{% /example %}}"}}
{{- end }}
{{"{{% /examples %}}"}}
`
	res, ok := pkgSpec.Resources[resourceName]
	if !ok {
		return fmt.Errorf("missing resource from schema: %s", resourceName)
	}

	t, err := template.New("examples").Funcs(template.FuncMap{
		"beginLanguage": func(lang interface{}) string {
			l := fmt.Sprintf("%s", lang)
			switch l {
			case "nodejs":
				l = "typescript"
			case "dotnet":
				l = "csharp"
			}
			return fmt.Sprintf("```%s", l)
		},
		"endLanguage": func() string {
			return "```"
		},
	}).Parse(tmpl)
	if err != nil {
		return err
	}
	b := strings.Builder{}
	if err = t.Execute(&b, examplesRenderData); err != nil {
		return err
	}
	res.Description += b.String()
	pkgSpec.Resources[resourceName] = res
	return nil
}

var excludeRegexes []*regexp.Regexp

func ShouldExclude(pulumiToken string) bool {
	for _, re := range excludeRegexes {
		if re.MatchString(pulumiToken) {
			return true
		}
	}

	return false
}

// InMemoryPackageLoader prevents having to fetch the schema from
// the provider every time which significantly speeds up codegen.
func InMemoryPackageLoader(pkgs map[string]*schema.Package) schema.Loader {
	return &inMemoryLoader{pkgs: pkgs}
}

type inMemoryLoader struct {
	pkgs map[string]*schema.Package
}

func (l *inMemoryLoader) LoadPackage(pkg string, _ *semver.Version) (*schema.Package, error) {
	if p, ok := l.pkgs[pkg]; ok {
		return p, nil
	}

	return nil, errors.Errorf("package %s not found in the in-memory map", pkg)
}

// camel replaces the first contiguous string of upper case runes in the given string with its lower-case equivalent.
func camel(s string) string {
	c, sz := utf8.DecodeRuneInString(s)
	if sz == 0 || unicode.IsLower(c) {
		return s
	}

	// The first rune is not lowercase. Iterate until we find a rune that is.
	var word []rune
	for {
		s = s[sz:]

		n, nsz := utf8.DecodeRuneInString(s)
		if nsz == 0 {
			word = append(word, unicode.ToLower(c))
			return string(word)
		}
		if unicode.IsLower(n) {
			if len(word) == 0 {
				c = unicode.ToLower(c)
			}
			word = append(word, c)
			return string(word) + s
		}
		c, sz, word = n, nsz, append(word, unicode.ToLower(c))
	}
}

// PlainLit returns an unquoted string literal expression.
func PlainLit(v string) *model.LiteralValueExpression {
	return &model.LiteralValueExpression{Value: cty.StringVal(v)}
}

// QuotedLit returns a quoted string literal expression.
func QuotedLit(v string) *model.TemplateExpression {
	return &model.TemplateExpression{Parts: []model.Expression{PlainLit(v)}}
}

// ObjectConsItem returns a new ObjectConsItem with the given key and value.
func ObjectConsItem(key string, value model.Expression) model.ObjectConsItem {
	var keyExpr model.Expression = PlainLit(key)
	if !hclsyntax.ValidIdentifier(key) {
		keyExpr = QuotedLit(key)
	}
	return model.ObjectConsItem{
		Key:   keyExpr,
		Value: value,
	}
}

// null represents PCL's builtin `null` variable
var null = &model.Variable{
	Name:         "null",
	VariableType: model.NoneType,
}

// RenderValue renders an object that represents a json value into PCL. Most nodes are rendered as one
// would expect (e.g. sequences -> tuple construction, maps -> object construction, etc.).
func RenderValue(node interface{}) (model.Expression, error) {
	if node == nil {
		return model.VariableReference(null), nil
	}

	typ := reflect.TypeOf(node)
	val := reflect.ValueOf(node)
	kind := typ.Kind()
	switch kind {
	case reflect.Slice:
		var expressions []model.Expression
		for i := 0; i < val.Len(); i++ {
			e, err := RenderValue(val.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			expressions = append(expressions, e)
		}
		return &model.TupleConsExpression{
			Expressions: expressions,
		}, nil
	case reflect.Bool:
		b := reflect.ValueOf(node).Bool()
		return &model.LiteralValueExpression{
			Value: cty.BoolVal(b),
		}, nil
	case reflect.Float32, reflect.Float64:
		return &model.LiteralValueExpression{
			Value: cty.NumberFloatVal(val.Float()),
		}, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &model.LiteralValueExpression{Value: cty.NumberIntVal(val.Int())}, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &model.LiteralValueExpression{Value: cty.NumberUIntVal(val.Uint())}, nil
	case reflect.String:
		return QuotedLit(val.String()), nil
	case reflect.Map:
		consItems := map[string]model.Expression{}
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			rendered, err := RenderValue(v.Interface())
			if err != nil {
				return nil, err
			}
			consItems[k.String()] = rendered
		}
		var items []model.ObjectConsItem
		for _, k := range codegen.SortedKeys(consItems) {
			items = append(items, ObjectConsItem(k, consItems[k]))
		}
		return &model.ObjectConsExpression{
			Items: items,
		}, nil
	case reflect.Ptr:
		nodeExpr, ok := node.(model.Expression)
		if !ok {
			// only expect model.Expression as the embedded interface
			panic(val)
		}
		return nodeExpr, nil
	default:
		contract.Failf("unexpected type %T", node)
		return nil, nil
	}
}

type programGenFn func(*hcl2.Program) (map[string][]byte, hcl.Diagnostics, error)

func generateExamplePrograms(token string, example Sample, body *model.Body, languages []string,
	bindOptions ...hcl2.BindOption) (languageToExampleProgram, error) {
	programBody := fmt.Sprintf("%v", body)
	fmt.Println(programBody)
	parser := syntax.NewParser()
	if err := parser.ParseFile(strings.NewReader(programBody), "program.pp"); err != nil {
		return nil, fmt.Errorf("failed to parse IR - token: %s: %v", token, err)
	}
	if parser.Diagnostics.HasErrors() {
		fmt.Println(programBody)
		err := parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	program, diags, err := hcl2.BindProgram(parser.Files, bindOptions...)
	if err != nil {
		log.Fatalf("failed to bind program: %v", err)
	}
	if diags.HasErrors() {
		log.Print(programBody)
		err := program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}

	languageExample := languageToExampleProgram{}

	for _, lang := range languages {
		var files map[string][]byte

		switch lang {
		case "dotnet":
			files, err = recoverableProgramGen(program, dotnet.GenerateProgram)
		case "go":
			files, err = recoverableProgramGen(program, gogen.GenerateProgram)
		case "nodejs":
			files, err = recoverableProgramGen(program, nodejs.GenerateProgram)
		case "python":
			files, err = recoverableProgramGen(program, python.GenerateProgram)
		default:
			continue
		}
		if err != nil {
			log.Printf("Program generation failed for language: %q for %q, continuing", lang, token)
			continue
		}

		buf := strings.Builder{}
		for _, f := range files {
			_, err := buf.Write(f)
			if err != nil {
				return nil, err
			}
		}
		languageExample[language(lang)] = programText(buf.String())
		fmt.Printf("Generated %s equivalent for %s\n", lang, token)
		fmt.Printf("%s", buf.String())
	}

	return languageExample, nil
}

func recoverableProgramGen(program *hcl2.Program, fn programGenFn) (files map[string][]byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered during generation: %v", r)
		}
	}()

	var d hcl.Diagnostics
	files, d, err = fn(program)

	if err != nil {
		return nil, err
	}
	if d.HasErrors() {
		err := program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(d)
		if err != nil {
			log.Printf("failed to write diagnostics: %v", err)
		}
	}
	return
}
