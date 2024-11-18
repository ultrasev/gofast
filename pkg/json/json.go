package json

import (
	"encoding/json"
	"io"
)

// Encode converts a value to JSON bytes
func Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode parses JSON bytes into a value
func Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// DecodeReader parses JSON from a reader into a value
func DecodeReader(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// IsValidJSON checks if a byte slice contains valid JSON
func IsValidJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// Pretty returns formatted JSON string with indentation
func Pretty(data []byte) ([]byte, error) {
	var temp interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return nil, err
	}
	return json.MarshalIndent(temp, "", "  ")
}

// Flatten converts nested JSON structure to flat key-value pairs
func Flatten(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	flattenRecursive(data, "", result)
	return result
}

// Helper function for Flatten
func flattenRecursive(data map[string]interface{}, prefix string, result map[string]interface{}) {
	for k, v := range data {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}

		switch child := v.(type) {
		case map[string]interface{}:
			flattenRecursive(child, key, result)
		default:
			result[key] = v
		}
	}
}
