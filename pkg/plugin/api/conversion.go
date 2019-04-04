package api

import (
	"encoding/json"

	cacheutil "github.com/heptio/developer-dash/internal/cache/util"
	"github.com/heptio/developer-dash/pkg/plugin/api/proto"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
)

func convertFromKey(in cacheutil.Key) (*proto.KeyRequest, error) {
	return &proto.KeyRequest{
		Namespace:  in.Namespace,
		ApiVersion: in.APIVersion,
		Kind:       in.Kind,
		Name:       in.Name,
	}, nil
}

func convertToKey(in *proto.KeyRequest) (cacheutil.Key, error) {
	if in == nil {
		return cacheutil.Key{}, errors.New("key request is nil")
	}

	matchLabels := labels.Set{}

	value := in.GetLabelSelector()
	if value != nil {
		if err := json.Unmarshal(value.Value, &matchLabels); err != nil {
			return cacheutil.Key{}, errors.Wrap(err, "unmarshal label selector")
		}
	}

	key := cacheutil.Key{
		Namespace:  in.Namespace,
		APIVersion: in.ApiVersion,
		Kind:       in.Kind,
		Name:       in.Name,
	}

	if len(matchLabels) > 0 {
		key.Selector = &matchLabels
	}

	return key, nil
}

func convertFromObjects(in []*unstructured.Unstructured) ([][]byte, error) {
	var out [][]byte

	for _, object := range in {
		data, err := convertFromObject(object)
		if err != nil {
			return nil, err
		}

		out = append(out, data)
	}

	return out, nil
}

func convertFromObject(in *unstructured.Unstructured) ([]byte, error) {
	return json.Marshal(in)
}

func convertToObjects(in [][]byte) ([]*unstructured.Unstructured, error) {
	var out []*unstructured.Unstructured

	for _, data := range in {
		object, err := convertToObject(data)
		if err != nil {
			return nil, err
		}

		out = append(out, object)
	}

	return out, nil
}

func convertToObject(in []byte) (*unstructured.Unstructured, error) {
	if in == nil {
		return nil, errors.New("can't convert nil object")
	}

	object := unstructured.Unstructured{}
	err := json.Unmarshal(in, &object)
	if err != nil {
		return nil, err
	}

	return &object, nil
}

func convertToPortForwardRequest(in *proto.PortForwardRequest) (*PortForwardRequest, error) {
	if in == nil {
		return nil, errors.New("can't convert nil object")
	}

	port := in.PortNumber
	if port > 0xFFFF {
		return nil, errors.Errorf("port number must be a uint32; it was: %d", port)
	}

	return &PortForwardRequest{
		Namespace: in.Namespace,
		PodName:   in.PodName,
		Port:      uint16(port),
	}, nil
}
