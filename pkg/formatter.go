package tinj

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/go-bongo/go-dotaccess"
)

const (
	// FieldExpression to capture field names
	FieldExpression = `\((\w*\|?\w*[\.?\w*]*)\)`
)

type LineFormatter struct {
	Fields []*Field
}

func CreateLineFormatter(fields []*Field) *LineFormatter {
	return &LineFormatter{Fields: fields}
}

func (l *LineFormatter) Print(line []rune) {
	var dict map[string]interface{}

	if !isJSON(string(line)) {
		fmt.Print(string(line))
		return
	}

	json.Unmarshal([]byte(string(line)), &dict)

	for _, field := range l.Fields {
		value, _ := dotaccess.Get(dict, field.Key)
		if value != nil {
			field.Print(value)
			fmt.Print(" | ")
		}
	}
	fmt.Print("\n")
}

// DeconstructFormat creates a line formatter from format string
func DeconstructFormat(format string) *LineFormatter {
	var fields []*Field

	r, _ := regexp.Compile(FieldExpression)

	for _, fieldInfo := range r.FindAllString(format, -1) {
		field := CreateField(fieldInfo)
		fields = append(fields, field)
	}
	return CreateLineFormatter(fields)
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
