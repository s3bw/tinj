/*
- Provide format

`[(cyan|service)]|(blue|severity)|"(message)" |(exc_info)|"(red|httpRequest.status)`
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/go-bongo/go-dotaccess"
)

// NewLine rune
const NewLine = '\n'

// FieldExpression to capture field names
const FieldExpression = `\((\w*\|?\w*[\.?\w*]*)\)`

func DeconstructFormat(format string) *LineFormatter {
	var fields []*Field
	var colour, key string

	r, _ := regexp.Compile(FieldExpression)

	for _, field := range r.FindAllString(format, -1) {
		colour = "white"
		key = field[1 : len(field)-1]

		if strings.Contains(key, "|") {
			arr := strings.Split(key, "|")
			colour, key = arr[0], arr[1]
		}

		newField := CreateField(key, colour)
		fields = append(fields, newField)
	}
	return CreateLineFormatter(fields)
}

type Field struct {
	Key    string
	Colour color.Attribute
}

func CreateField(field, colour string) *Field {
	var colourTable = map[string]color.Attribute{
		"white": color.FgWhite,
		"cyan":  color.FgCyan,
		"red":   color.FgRed,
		"blue":  color.FgBlue,
	}
	return &Field{
		Key:    field,
		Colour: colourTable[colour],
	}
}

type LineFormatter struct {
	Fields []*Field
}

func CreateLineFormatter(fields []*Field) *LineFormatter {
	return &LineFormatter{Fields: fields}
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func (l *LineFormatter) Construct(line []rune) {
	var dict map[string]interface{}

	if !isJSON(string(line)) {
		fmt.Print(string(line))
		return
	}

	json.Unmarshal([]byte(string(line)), &dict)

	for _, field := range l.Fields {
		outLine, _ := dotaccess.Get(dict, field.Key)
		if outLine != nil {
			color.Set(field.Colour)
			fmt.Print(outLine)
			fmt.Print(" | ")
			color.Unset()
		}
	}
	fmt.Print("\n")
}

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	// Help Text
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: cat file.json | tinj")
		return
	}
	lineFormatter := DeconstructFormat(`[(cyan|service)]|(blue|severity)|(red|httpRequest.status)|"(message)"|(exc_info)|`)

	var line []rune
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		line = append(line, input)

		if input == NewLine {
			lineFormatter.Construct(line)
			line = nil
		}
	}
}
