package tinj

import (
	"encoding/json"
	"fmt"

	"github.com/go-bongo/go-dotaccess"
)

// Add the separator to the line formatter
type LineFormatter struct {
	Fields []*Field
}

func CreateLineFormatter(fields []*Field) *LineFormatter {
	return &LineFormatter{Fields: fields}
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
			fmt.Print(" | ")
		}
	}
	fmt.Print("\n")
}
