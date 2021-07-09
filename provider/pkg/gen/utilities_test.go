// Copyright 2016-2021, Pulumi Corporation.

package gen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var apiParamToSdkNames = map[string]string{
	"projectsId":               "project",
	"locationId":               "location",
	"managedZonesId":           "managedZone",
	"gameServerDeploymentsId":  "gameServerDeploymentId",
	"schemaId":                 "schemaId",
	"certificateAuthoritiesId": "certificateAuthorityId",
	"sourceContents":           "sourceContents",
	"dataId":                   "dataId",
}

func TestApiParamNameToSdkName(t *testing.T) {
	for name, expected := range apiParamToSdkNames {
		actual := apiParamNameToSdkName(name)
		assert.Equal(t, expected, actual)
	}
}

var apiPropToSdkNames = map[string]string{
	"projectId":         "project",
	"locationId":        "location",
	"managedZoneId":     "managedZone",
	"ApiName":           "apiName",
	"IPAddress":         "ipAddress",
	"custom_attributes": "customAttributes",
}

func TestApiPropNameToSdkName(t *testing.T) {
	for name, expected := range apiPropToSdkNames {
		actual := apiPropNameToSdkName(name)
		assert.Equal(t, expected, actual)
	}
}
