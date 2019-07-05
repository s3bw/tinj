package tinj

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	// ColourSeparator separates colour and field name
	ColourSeparator = "|"
	// DefaultColour if colour isn't provided
	DefaultColour = "white"
)

// Field identifies a field in a JSON to format into a log
type Field struct {
	// Key specifying the value to find in the json
	Key string
	// Colour specifying the colour to output the value
	Colour color.Attribute
}

func CreateField(fieldSpec string) *Field {
	var colour, key string
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

	colour = DefaultColour
	key = fieldSpec[1 : len(fieldSpec)-1]
	if strings.Contains(key, ColourSeparator) {
		arr := strings.Split(key, ColourSeparator)
		colour, key = arr[0], arr[1]
	}

	return &Field{
		Key:    key,
		Colour: colourTable[colour],
	}
}

func (f *Field) Print(line interface{}) {
	color.Set(f.Colour)
	fmt.Print(line)
	color.Unset()
}
