// Package parser marshals and unmarshals protocol buffer messages in YAML format.
package parser

import (
	"fmt"

	"github.com/austinthao5/golang_proto_test/config/deploymentconfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/fileio"
	"github.com/austinthao5/golang_proto_test/internal/validate"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"
)

var nonStrict = protojson.UnmarshalOptions{DiscardUnknown: true}
var strict = protojson.UnmarshalOptions{DiscardUnknown: false}

// Marshal writes the given proto.Message in YAML format.
func Marshal(m proto.Message) ([]byte, error) {
	json, err := protojson.Marshal(m)
	if err != nil {
		return nil, err
	}
	return yaml.JSONToYAML(json)
}

// Unmarshal reads the given []byte into the given proto.Message, discarding
// any unknown fields in the input.
func Unmarshal(b []byte, m proto.Message) error {
	json, err := yaml.YAMLToJSON(b)
	if err != nil {
		return err
	}
	return nonStrict.Unmarshal(json, m)
}

// UnmarshalStrict reads the given []byte into the given proto.Message. If there
// are any unknown fields in the input, an error is returned.
func UnmarshalStrict(b []byte, m proto.Message) error {
	json, err := yaml.YAMLToJSON(b)
	if err != nil {
		return err
	}
	return strict.Unmarshal(json, m)
}

// ParseHalConfig reads the YAML file at halPath parses it into a *config.Hal.
// The config.Hal is validated; if there are any errors, they will be returned
// in error and *config.Hal will be nil.
// func ParseHalConfig(halPath string) (*config.Hal, error) {
func ParseHalConfig(halPath string) (*providers.AppEngine, error) {
	data, err := fileio.ReadFile(halPath)
	if err != nil {
		return nil, err
	}

	// Debug
	// fmt.Println("===RAW hal data START===\n" + string(data) + "\n===RAW hal data END===")

	// hal := &config.Hal{}
	hal := &providers.AppEngine{}
	if err := Unmarshal(data, hal); err != nil {
		return nil, fmt.Errorf("unable to unmarshal %q: %v", halPath, err)
	}

	if err := validate.HalConfig(hal); err != nil {
		return nil, fmt.Errorf("unable to validate %q: %v", halPath, err)
	}

	return hal, nil
}