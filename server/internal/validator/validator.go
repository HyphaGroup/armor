package validator

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Validator struct {
	schemas map[string]*jsonschema.Schema
}

type ValidationError struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

func New(schemasDir string) (*Validator, error) {
	v := &Validator{
		schemas: make(map[string]*jsonschema.Schema),
	}

	sectionFiles := map[string]string{
		"mission":     "mission.schema.json",
		"assets":      "assets.schema.json",
		"adversaries": "adversaries.schema.json",
		"threats":     "threats.schema.json",
		"risks":       "risks.schema.json",
		"mitigations": "mitigations.schema.json",
	}

	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft7

	for section, filename := range sectionFiles {
		path := filepath.Join(schemasDir, filename)

		schemaURL := "file://" + path
		schema, err := compiler.Compile(schemaURL)
		if err != nil {
			return nil, fmt.Errorf("failed to compile schema %s: %w", filename, err)
		}

		v.schemas[section] = schema
	}

	return v, nil
}

func (v *Validator) Validate(section, data string) ([]ValidationError, error) {
	schema, ok := v.schemas[section]
	if !ok {
		return nil, fmt.Errorf("unknown section: %s", section)
	}

	var jsonData interface{}
	if err := json.Unmarshal([]byte(data), &jsonData); err != nil {
		return []ValidationError{{
			Path:    "$",
			Message: fmt.Sprintf("invalid JSON: %s", err.Error()),
		}}, nil
	}

	err := schema.Validate(jsonData)
	if err == nil {
		return nil, nil
	}

	validationErr, ok := err.(*jsonschema.ValidationError)
	if !ok {
		return nil, fmt.Errorf("unexpected validation error type: %w", err)
	}

	var errors []ValidationError
	collectErrors(validationErr, &errors)

	return errors, nil
}

func collectErrors(err *jsonschema.ValidationError, errors *[]ValidationError) {
	if len(err.Causes) == 0 {
		*errors = append(*errors, ValidationError{
			Path:    err.InstanceLocation,
			Message: err.Message,
		})
		return
	}

	for _, cause := range err.Causes {
		collectErrors(cause, errors)
	}
}

func (v *Validator) HasSchema(section string) bool {
	_, ok := v.schemas[section]
	return ok
}
