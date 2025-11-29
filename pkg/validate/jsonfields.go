package validate

import (
	"encoding/json"
	"fmt"
)

func RequiredFields(data []byte, fields ...string) error {
	var raw map[string]any

	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("validate fields: %w", err)
	}

	for _, field := range fields {
		if _, ok := raw[field]; !ok {
			return fmt.Errorf("validate fields: missing required field %q", field)
		}
	}

	return nil
}
