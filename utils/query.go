package utils

import (
	"strings"
)

// QueryParam represents a key-value pair for URL query parameters.
type QueryParam struct {
	Key   string
	Value string
}

// NewQueryParam creates a new QueryParam instance with the provided attribute and value.
func NewQueryParam(attr string, v interface{}) QueryParam {
	switch value := v.(type) {
	case string:
		return QueryParam{Key: attr, Value: value}
	case []string:
		return QueryParam{Key: attr, Value: strings.Join(value, ",")}
	default:
		panic("unsupported value query params type")
	}
}
