package selector

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// ParseRule unmarshals YAML data into a Rule and validates it.
func ParseRule(data []byte) (*Rule, error) {
	var r Rule
	if err := yaml.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("failed to unmarshal selector: %w", err)
	}

	// Set default operand if not specified
	if r.Operand == "" {
		r.Operand = "any"
	}

	// Normalize operand
	switch strings.ToLower(r.Operand) {
	case "and", "all":
		r.Operand = "and"
	case "or", "any":
		r.Operand = "or"
	case "not", "none":
		r.Operand = "not"
	default:
		return nil, fmt.Errorf("invalid operand: %s", r.Operand)
	}

	// Initialize Tags map if nil
	if r.Tags == nil && r.Type == "" {
		r.Tags = make(map[string]string)
	}

	// Validation
	switch r.Type {
	case "header", "jwt":
		if r.Key == "" || len(r.Values) == 0 {
			return nil, fmt.Errorf("type %s requires key and values", r.Type)
		}
	case "input.model":
		if len(r.Values) == 0 {
			return nil, fmt.Errorf("type input.model requires values")
		}
	case "":
		if len(r.Tags) == 0 {
			return nil, nil // No type and no tags means no rule
		}
	default:
		return nil, fmt.Errorf("unsupported selector type: %s", r.Type)
	}

	return &r, nil
}
