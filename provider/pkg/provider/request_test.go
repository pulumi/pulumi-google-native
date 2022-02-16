// Copyright 2016-2021, Pulumi Corporation.

package provider

import (
	"testing"

	"github.com/pulumi/pulumi-google-native/provider/pkg/gen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestNetworkPatchRequest(t *testing.T) {
	_, meta, _ := gen.PulumiSchema("../../../discovery")
	res := meta.Resources["google-native:compute/v1:Subnetwork"]
	// Passes if this is uncommented
	// delete(res.UpdateProperties, "region")
	p := googleCloudProvider{}
	inputs := resource.PropertyMap{
		"region": resource.PropertyValue{
			V: "region-a",
		},
	}
	state := resource.PropertyMap{}
	actual := p.prepareAPIInputs(inputs, state, res.UpdateProperties)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, actual)
}
