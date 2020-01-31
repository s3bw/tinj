package tinj

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"os"

	"github.com/fatih/color"
	"github.com/go-bongo/go-dotaccess"
)

// Style of logs to parse into readable lines
type Style int

const (
	// Tail style is ordinary style
	Tail Style = iota
	// Stern multi-pod logs https://github.com/wercker/stern
	Stern
	// Compose is used for docker-compose
	Compose
)

// LineFormatter turns a JSON input into a line for the command line
type LineFormatter struct {
	Fields    []*Field
	Separator string
	Style     Style
}

var colorList = []*color.Color{
	color.New(color.FgHiCyan),
	color.New(color.FgHiGreen),
	color.New(color.FgHiMagenta),
	color.New(color.FgHiYellow),
	color.New(color.FgHiBlue),
	color.New(color.FgHiRed),
}

func CreateLineFormatter(fields []*Field, separator string, style Style) *LineFormatter {
	return &LineFormatter{
		Fields:    fields,
		Separator: separator,
		Style:     style,
	}
}

// ReadStdin streams lines from stdin
func ReadStdin(format, separator string, style Style) {
	fields := ConstructFields(format)
	lineFormatter := CreateLineFormatter(fields, separator, style)

	stdin := bufio.NewReader(os.Stdin)
	for {
		nextLine, err := stdin.ReadBytes('\n')
		if err != nil {
			// Exit? Exit on interrupt!
			break
		}
		lineFormatter.Print(nextLine)
		nextLine = nil
	}
}

// Print line with new format
func (l *LineFormatter) Print(line []byte) {
	prefix, line := l.normPrefix(line)
	if prefix != nil {
		colour := determineColor(prefix)
		fmt.Printf("%v", colour(string(prefix)))
	}

	var dict map[string]interface{}
	err := json.Unmarshal(line, &dict)
	if err != nil {
		// If there is an error in parsing the log
		// we still want to output the log to stdout
		fmt.Println(string(line))
		return
	}

	for i, field := range l.Fields {
		value, _ := dotaccess.Get(dict, field.Key)
		if value != nil {
			if i != 0 {
				fmt.Print(l.Separator)
			}
			field.Print(value)
		}
	}
	fmt.Print("\n")
}

func (l *LineFormatter) normPrefix(line []byte) ([]byte, []byte) {
	if (len(line) > 0) && (line[0] != '{') {
		switch l.Style {
		case Tail:
			return nil, line
		default:
			for i, b := range line {
				if '{' == b {
					return line[:i], line[i:]
				}
			}
		}
	}
	return nil, line
}

func determineColor(prefix []byte) func(...interface{}) string {
	hash := fnv.New32()
	hash.Write([]byte(prefix))

	idx := hash.Sum32() % uint32(len(colorList))
	return colorList[idx].SprintFunc()
}
