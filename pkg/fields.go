package tinj

import (
	"fmt"

	"github.com/fatih/color"
)

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

func (f *Field) Print(line interface{}) {
	color.Set(f.Colour)
	fmt.Print(line)
	color.Unset()
}
