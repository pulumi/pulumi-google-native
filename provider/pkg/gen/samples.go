package gen

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"

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
func ExtractSamples(baseFs fs.FS, pkgSpec *schema.PackageSpec) (map[string][]Sample, error) {
	const basePattern = "services/google/**/samples/*.yaml"

	serviceGrouping, err := serviceGrouping(pkgSpec)
	if err != nil {
		return nil, err
	}

	samples := map[string][]Sample{}
	if err := dstar.GlobWalk(baseFs, basePattern, func(path string, d fs.DirEntry) error {
		s, err := parseSample(baseFs, path, serviceGrouping)
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

func parseSample(fs fs.FS, path string, grouping map[string]tokenGrouping) (map[string][]Sample, error) {
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
	resourceName := ToUpperCamel(splitResource[1])

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
	sample := Sample{
		Description: metadata.Description,
		Body:        raw,
	}
	for _, v := range metadata.Versions {
		toks := findTokens(grp, serviceName, resourceName, v)
		for _, tok := range toks {
			samples[tok] = []Sample{sample}
		}
	}

	return samples, nil
}
