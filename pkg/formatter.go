package tinj

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-bongo/go-dotaccess"
)

const (
	// FieldExpression to capture field names
	FieldExpression = `\((\w*\|?\w*[\.?\w*]*)\)`
	// ColourSeparator separates colour and field name
	ColourSeparator = "|"
	// DefaultColour if colour isn't provided
	DefaultColour = "white"
)

type LineFormatter struct {
	Fields []*Field
}

func CreateLineFormatter(fields []*Field) *LineFormatter {
	return &LineFormatter{Fields: fields}
}

func (l *LineFormatter) Format(line []rune) {
	var dict map[string]interface{}

	if !isJSON(string(line)) {
		fmt.Print(string(line))
		return
	}

	json.Unmarshal([]byte(string(line)), &dict)

	for _, field := range l.Fields {
		outLine, _ := dotaccess.Get(dict, field.Key)
		if outLine != nil {
			field.Print(outLine)
			fmt.Print(" | ")
		}
	}
	fmt.Print("\n")
}

// DeconstructFormat creates a line formatter from format string
func DeconstructFormat(format string) *LineFormatter {
	var fields []*Field
	var colour, key string

	r, _ := regexp.Compile(FieldExpression)

	for _, field := range r.FindAllString(format, -1) {
		colour = DefaultColour
		key = field[1 : len(field)-1]

		if strings.Contains(key, ColourSeparator) {
			arr := strings.Split(key, ColourSeparator)
			colour, key = arr[0], arr[1]
		}

		newField := CreateField(key, colour)
		fields = append(fields, newField)
	}
	return CreateLineFormatter(fields)
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
