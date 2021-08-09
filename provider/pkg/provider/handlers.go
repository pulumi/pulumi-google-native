// Copyright 2021, Pulumi Corporation.

package provider

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"

	containerv1 "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container"
	containerv1beta1 "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container/beta"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

type updateHandler interface {
	Update(resource resources.CloudAPIResource, inputs, oldState resource.PropertyMap) (map[string]interface{}, error)
}

type updateHandlerFunc func(resource resources.CloudAPIResource, inputs, oldState resource.PropertyMap) (map[string]interface{}, error)

func (u updateHandlerFunc) Update(resource resources.CloudAPIResource, inputs, oldState resource.PropertyMap) (map[string]interface{}, error) {
	return u(resource, inputs, oldState)
}

func (p *googleCloudProvider) updateClusterV1DCL(req *pulumirpc.UpdateRequest) updateHandlerFunc {
	return func(res resources.CloudAPIResource, inputs, oldState resource.PropertyMap) (map[string]interface{}, error) {
		body := p.prepareAPIInputs(inputs, oldState, res.CreateProperties, true)

		var cluster containerv1.Cluster
		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Result:  &cluster,
			TagName: "json",
		})
		if err != nil {
			return nil, err
		}
		logging.V(9).Infof("Desired raw state: %+v", body)
		logging.V(9).Infof("Old state: %+v", oldState)

		if err := decoder.Decode(body["cluster"]); err != nil {
			return nil, err
		}

		urn := resource.URN(req.GetUrn())
		label := fmt.Sprintf("%s.Diff(%s)", p.name, urn)
		oldStateRaw, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
			Label:        fmt.Sprintf("%s.oldState", label),
			KeepUnknowns: true,
			KeepSecrets:  true,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "failed because malformed resource inputs")
		}

		// Extract old inputs from the `__inputs` field of the old state.
		oldInputs := parseCheckpointObject(oldStateRaw)
		if cluster.Location == nil {
			if oldInputs.HasValue("location") {
				location := oldInputs["location"].StringValue()
				cluster.Location = &location
			}
		}
		if cluster.Project == nil {
			if oldInputs.HasValue("project") {
				project := oldInputs["project"].StringValue()
				cluster.Project = &project
			}
		}

		ctx := context.Background()
		client, err := p.client.HttpClient(ctx)
		if err != nil {
			return nil, err
		}
		// TODO: figure out logging.
		containerClient := containerv1.NewClient(dcl.NewConfig(dcl.WithHTTPClient(client)))

		logging.V(9).Infof("Desired raw state of cluster: %+v", cluster)
		resp, err := containerClient.ApplyCluster(ctx, &cluster)
		if err != nil {
			logging.V(1).Infof("%+v", err)
			return nil, err
		}
		result := map[string]interface{}{}
		decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Result:  &result,
			TagName: "json",
		})
		if err != nil {
			return nil, err
		}

		if err = decoder.Decode(resp); err != nil {
			return nil, err
		}
		return result, nil
	}
}

func (p *googleCloudProvider) updateClusterV1Beta1DCL(req *pulumirpc.UpdateRequest) updateHandlerFunc {
	return func(res resources.CloudAPIResource, inputs, oldState resource.PropertyMap) (map[string]interface{}, error) {
		body := p.prepareAPIInputs(inputs, oldState, res.CreateProperties, true)
		var cluster containerv1beta1.Cluster
		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Result:  &cluster,
			TagName: "json",
		})
		if err != nil {
			return nil, err
		}
		logging.V(9).Infof("Desired raw state: %+v", body)
		if err := decoder.Decode(body["cluster"]); err != nil {
			return nil, err
		}
		ctx := context.Background()
		client, err := p.client.HttpClient(ctx)
		if err != nil {
			return nil, err
		}
		// TODO: figure out logging.
		containerClient := containerv1beta1.NewClient(dcl.NewConfig(dcl.WithHTTPClient(client)))

		logging.V(9).Infof("Desired raw state of cluster: %+v", cluster)
		resp, err := containerClient.ApplyCluster(ctx, &cluster)
		if err != nil {
			return nil, err
		}
		result := map[string]interface{}{}
		decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Result:  &result,
			TagName: "json",
		})
		if err != nil {
			return nil, err
		}

		if err = decoder.Decode(resp); err != nil {
			return nil, err
		}
		return result, nil
	}
}

func (p *googleCloudProvider) updateHandler(resourceTok string, req *pulumirpc.UpdateRequest) (updateHandler, bool) {
	logging.V(9).Infof("UpdateHandler for %q", resourceTok)
	switch resourceTok {
	case "google-native:container/v1:Cluster":
		return p.updateClusterV1DCL(req), true
	case "google-native:container/v1beta1:Cluster":
		return p.updateClusterV1Beta1DCL(req), true
	}
	return nil, false
}
