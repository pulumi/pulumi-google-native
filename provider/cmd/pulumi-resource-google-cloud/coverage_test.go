package main

import (
	"encoding/json"
	"fmt"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestCoverage(t *testing.T) {
	f, err := os.Open("schema.json")
	require.NoError(t, err)
	j := json.NewDecoder(f)
	native := schema.PackageSpec{}
	require.NoError(t, j.Decode(&native))
	f.Close()
	resourceToNativeTok := map[string]codegen.StringSet{}
	for tok := range native.Resources {
		sp := strings.Split(tok, ":")
		res := sp[len(sp) - 1]
		if _, ok := resourceToNativeTok[res]; !ok {
			resourceToNativeTok[res] = codegen.NewStringSet()
		}
		resourceToNativeTok[res].Add(tok)
	}

	resp, err := http.Get("https://raw.githubusercontent.com/pulumi/pulumi-gcp/master/provider/cmd/pulumi-resource-gcp/schema.json")
	require.NoError(t, err)
	j = json.NewDecoder(resp.Body)
	tf := schema.PackageSpec{}
	require.NoError(t, j.Decode(&tf))
	resourceToTfTok := map[string]codegen.StringSet{}
	for tok := range tf.Resources {
		sp := strings.Split(tok, ":")
		res := sp[len(sp) -1]
		if _, ok := resourceToTfTok[res]; !ok {
			resourceToTfTok[res] = codegen.NewStringSet()
		}
		resourceToTfTok[res].Add(tok)
	}

	nativeKeys := codegen.SortedKeys(resourceToNativeTok)
	tfKeys := codegen.SortedKeys(resourceToTfTok)

	var nativeButNotTf, tfButNotNative, inBoth []string
	var nativeIndex, tfIndex int
	var currNative, currTf string
	for {
		if nativeIndex == len(nativeKeys) && tfIndex == len(tfKeys) {
			break
		}
		if nativeIndex < len(nativeKeys) {
			currNative = nativeKeys[nativeIndex]
		}
		if tfIndex < len(tfKeys) {
			currTf = tfKeys[tfIndex]
		}
		if currNative < currTf {
			nativeButNotTf = append(nativeButNotTf, currNative)
			nativeIndex++
		} else if currTf < currNative {
			tfButNotNative = append(tfButNotNative, currTf)
			tfIndex++
		} else {
			inBoth = append(inBoth, currTf)
			tfIndex++
			nativeIndex++
		}
	}

	var fromTF, fromNative int
	for _, k := range inBoth {
		fromTF += len(resourceToTfTok[k])
		fromNative += len(resourceToNativeTok[k])
	}

	fmt.Printf("In both: %d\n", len(inBoth))
	fmt.Printf("Accounts for %d in native provider and %d in TF provider\n\n", fromNative, fromTF)

	missingInTF := codegen.NewStringSet()
	for _, k := range nativeButNotTf {
		for t := range resourceToNativeTok[k] {
			missingInTF.Add(t)
		}
	}

	missingInNative := codegen.NewStringSet()
	for _, k := range tfButNotNative {
		for t := range resourceToTfTok[k] {
			missingInNative.Add(t)
		}
	}

	fmt.Printf("Only in Native (%d) - unique (%d):\n", len(missingInTF), len(nativeButNotTf))
	for _, k := range missingInTF.SortedValues() {
		fmt.Println(k)
	}
	fmt.Println()

	fmt.Printf("Only in TF (%d) - unique (%d):\n", len(missingInNative), len(tfButNotNative))
	var iamRelated int
	for _, k := range missingInNative.SortedValues() {
		lowered := strings.ToLower(k)
		if !strings.HasSuffix(lowered, "iambinding") && !strings.HasSuffix(lowered, "iampolicy") && !strings.HasSuffix(lowered, "iammember") {
			fmt.Println(k)
		} else {
			iamRelated++
		}
	}
	fmt.Printf("%d are IAM related", iamRelated)
	fmt.Println()
}
