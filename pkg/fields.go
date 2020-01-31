package tinj

import (
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

// Field identifies a field in a JSON to format into a log
type Field struct {
	// Key specifying the value to find in the json
	Key string
	// Colour specifying the colour to output the value
	Colour color.Attribute
}

// ConstructFields parses output format from string
func ConstructFields(format string) []*Field {
	var fields []*Field

	r, _ := regexp.Compile(FieldExpression)
	for _, fieldInfo := range r.FindAllString(format, -1) {
		fieldKey, colour := SplitFieldInfo(fieldInfo)
		field := CreateField(fieldKey, colour)
		fields = append(fields, field)
	}
	return fields
}

// CreateField given a spec <colour>|<fieldName> or just <fieldName>
func CreateField(key, colour string) *Field {
	var colourTable = map[string]color.Attribute{
		"black":   color.FgBlack,
		"blue":    color.FgBlue,
		"cyan":    color.FgCyan,
		"green":   color.FgGreen,
		"magenta": color.FgMagenta,
		"red":     color.FgRed,
		"yellow":  color.FgYellow,
		"white":   color.FgWhite,
	}

	return &Field{
		Key:    key,
		Colour: colourTable[colour],
	}
}

// Print the value as per Field specification
func (field *Field) Print(text interface{}) {
	apply := color.New(field.Colour).SprintFunc()
	fmt.Printf("%s", apply(text))
}
