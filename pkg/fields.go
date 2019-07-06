package tinj

import (
	"fmt"

	"github.com/fatih/color"
)

// Field identifies a field in a JSON to format into a log
type Field struct {
	// Key specifying the value to find in the json
	Key string
	// Colour specifying the colour to output the value
	Colour color.Attribute
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
func (f *Field) Print(value interface{}) {
	color.Set(f.Colour)
	fmt.Print(value)
	color.Unset()
}
