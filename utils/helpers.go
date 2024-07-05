// Utility Functions

package utils

import (
	"encoding/json"
)

// GetStringOrDefault returns the value or a default value if the input value is nil.
func GetStringOrDefault(value *string, defaultValue string) string {
	if value == nil {
		return defaultValue
	}
	return *value
}

// GetBoolOrDefault returns the value or a default value if the input value is nil.
func GetBoolOrDefault(value *bool, defaultValue bool) bool {
	if value == nil {
		return defaultValue
	}
	return *value
}

// GetIntOrDefault returns the value or a default value if the input value is nil.
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

	if jsonString == nil || *jsonString == "" {
		jsonStr = defaultValue
	} else {
		jsonStr = *jsonString
	}

	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		// Log the error and return an empty JSON object
		return make(map[string]interface{})
	}

	return result
}

// ToJSONString converts a slice of strings to a JSON-encoded string.
func ToJSONString(slice []string) string {
	jsonData, err := json.Marshal(slice)
	if err != nil {
		return "[]"
	}
	return string(jsonData)
}
