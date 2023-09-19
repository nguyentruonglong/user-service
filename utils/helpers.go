// Utility Functions

package utils

import (
	"encoding/json"
)

// GetStringOrDefault returns the value or a default value if the input value is nil
func GetStringOrDefault(value *string, defaultValue string) string {
	if value == nil {
		return defaultValue
	}
	return *value
}

// GetBoolOrDefault returns the value or a default value if the input value is nil
func GetBoolOrDefault(value *bool, defaultValue bool) bool {
	if value == nil {
		return defaultValue
	}
	return *value
}

// GetIntOrDefault returns the value or a default value if the input value is nil
func GetIntOrDefault(value *int, defaultValue int) int {
	if value == nil {
		return defaultValue
	}
	return *value
}

// GetOrDefaultJSON takes a JSON string pointer and returns a map[string]interface{}.
// If the input JSON is empty or invalid, or the pointer is nil, it returns an empty JSON object.
func GetOrDefaultJSON(jsonString *string, defaultValue string) map[string]interface{} {
	var jsonStr string

	if jsonString == nil {
		jsonStr = defaultValue
	} else {
		jsonStr = *jsonString
		if jsonStr == "" {
			jsonStr = defaultValue
		}
	}

	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return make(map[string]interface{})
	}

	return result
}
