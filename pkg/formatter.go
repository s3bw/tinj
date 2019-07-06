package tinj

import (
	"encoding/json"
	"fmt"

	"github.com/go-bongo/go-dotaccess"
)

// LineFormatter turns a JSON input into a line for the command line
type LineFormatter struct {
	Fields    []*Field
	Separator string
}

func CreateLineFormatter(fields []*Field, separator string) *LineFormatter {
	return &LineFormatter{Fields: fields, Separator: separator}
}

func (l *LineFormatter) Print(line []byte) {
	var dict map[string]interface{}

	err := json.Unmarshal(line, &dict)

	if err != nil {
		fmt.Println(string(line))
		return
	}

	for _, field := range l.Fields {
		value, _ := dotaccess.Get(dict, field.Key)
		if value != nil {
			field.Print(value)
			fmt.Print(l.Separator)
		}
	}
	fmt.Print("\n")
}
