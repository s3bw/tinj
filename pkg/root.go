package tinj

import (
	"strings"
)

const (
	// DefaultColour if colour isn't provided
	DefaultColour = "white"
	// ColourSeparator separates colour and field name
	ColourSeparator = "|"
	// FieldExpression to capture field names
	// captures (colour|field), (field) and (deeply.nested.fields)
	FieldExpression = `\((\w*[\.?\w*]*\|?\w*)\)`
)

// SplitFieldInfo into fieldName and colour
func SplitFieldInfo(fieldInfo string) (string, string) {
	var colour, field string

	colour = DefaultColour
	field = fieldInfo[1 : len(fieldInfo)-1]

	// If colour is specified apply the colour otherwise use default
	if strings.Contains(field, ColourSeparator) {
		arr := strings.Split(field, ColourSeparator)
		field, colour = arr[0], arr[1]
	}
	return field, colour
}
