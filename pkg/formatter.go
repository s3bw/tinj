package tinj

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

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

// Print line with new format
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

// ReadStdin streams lines from stdin
func ReadStdin(format, separator string) {
	fields := ConstructFields(format)
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
