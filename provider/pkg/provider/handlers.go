// Copyright 2021, Pulumi Corporation.

package provider

import (
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

type updateHandler interface {
	Update(inputs resource.PropertyMap, oldState resource.PropertyMap) (map[string]interface{}, error)
}

type updateHandlerFunc func(inputs resource.PropertyMap, oldState resource.PropertyMap) (map[string]interface{}, error)

func (u updateHandlerFunc) Update(inputs resource.PropertyMap, oldState resource.PropertyMap) (map[string]interface{}, error) {
	return u(inputs, oldState)
}

func (p *googleCloudProvider) updateHandler(resourceTok string, req *rpc.UpdateRequest) updateHandler {
	logging.V(9).Infof("UpdateHandler for %q", resourceTok)
	switch resourceTok {
	case "google-native:container/v1:Cluster", "google-native:container/v1beta1:Cluster":
		return updateHandlerFunc(func(inputs, oldState resource.PropertyMap) (map[string]interface{}, error) {
			res, ok := p.resourceMap.Resources[resourceTok]
			if !ok {
				return nil, errors.Errorf("resource %q not found", resourceTok)
			}

			updateBody := p.prepareAPIInputs(inputs, oldState, res.UpdateProperties)

			// TODO: Cluster updates also include flattened members of the `networkConfig` map
			// which is hard to represent with the current metadata format in place (sdkName
			// can only be used to refer to fields at the top level).

			var splitBody []map[string]interface{}
			update, hasUpdate := updateBody["update"]
			updateMap, isMap := update.(map[string]interface{})
			if hasUpdate && isMap {
				for k, v := range updateMap {
					if strings.HasPrefix(k, "desired") {
						splitBody = append(splitBody, map[string]interface{}{k: v})
					}
				}
			}

			logging.V(9).Infof("updateBody: %+v", updateBody)
			logging.V(9).Infof("splitBody: %+v", splitBody)

			// Update operation on GKE clusters are to be performed one field at a time.
			mergedResp := map[string]interface{}{}
			for _, desiredBody := range splitBody {
				resp, err := p.updateRoundTrip(res.ResourceUrl(req.GetId()), res, map[string]interface{}{"update": desiredBody})
				if err != nil {
					return nil, err
				}
				if err = mergo.Merge(&mergedResp, resp); err != nil {
					return nil, err
				}
			}
			return mergedResp, nil
		})
	default:
		return updateHandlerFunc(func(inputs resource.PropertyMap, oldState resource.PropertyMap) (map[string]interface{}, error) {
			res, ok := p.resourceMap.Resources[resourceTok]
			if !ok {
				return nil, errors.Errorf("resource %q not found", resourceTok)
			}

			body := p.prepareAPIInputs(inputs, oldState, res.UpdateProperties)
			return p.updateRoundTrip(res.ResourceUrl(req.GetId()), res, body)
		})
	}
}

func (p *googleCloudProvider) updateRoundTrip(uri string, res resources.CloudAPIResource, body map[string]interface{}) (map[string]interface{}, error) {
	if strings.HasSuffix(uri, ":getIamPolicy") {
		uri = strings.ReplaceAll(uri, ":getIamPolicy", ":setIamPolicy")
	}

	op, err := p.client.RequestWithTimeout(res.UpdateVerb, uri, body, 0)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s: %q %+v", err, uri, body)
	}

	resp, err := p.waitForResourceOpCompletion(res.BaseUrl, op)
	if err != nil {
		return nil, errors.Wrapf(err, "waiting for completion")
	}
	return resp, nil
}
