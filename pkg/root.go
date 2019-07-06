package tinj

import (
	"bufio"
	"os"
	"regexp"
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

// ReadStdin streams lines from stdin
func ReadStdin(format, separator string) {
	fields := DeconstructFormat(format)
	lineFormatter := CreateLineFormatter(fields, separator)
	stdin := bufio.NewReader(os.Stdin)

	for {
		nextLine, _, err := stdin.ReadLine()
		if err != nil {
			break
		}
		lineFormatter.Print(nextLine)
		nextLine = nil
	}
}

// DeconstructFormat parses output format from string
func DeconstructFormat(format string) []*Field {
	var fields []*Field

	r, _ := regexp.Compile(FieldExpression)
	for _, fieldInfo := range r.FindAllString(format, -1) {
		fieldKey, colour := SplitFieldInfo(fieldInfo)
		field := CreateField(fieldKey, colour)
		fields = append(fields, field)
	}
	return fields
}

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
