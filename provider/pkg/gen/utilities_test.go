// Copyright 2016-2021, Pulumi Corporation.

package gen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var apiToSdkNames = map[string]string{
	"projectsId":               "project",
	"locationId":               "location",
	"managedZonesId":           "managedZone",
	"gameServerDeploymentsId":  "gameServerDeploymentId",
	"schemaId":                 "schemaId",
	"certificateAuthoritiesId": "certificateAuthorityId",
	"sourceContents":           "sourceContents",
	"dataId":                   "dataId",
}

func TestApiNameToSdkName(t *testing.T) {
	for name, expected := range apiToSdkNames {
		actual := apiNameToSdkName(name)
		assert.Equal(t, expected, actual)
	}
}
