package tinj

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
)

// Metric identifies a field in a JSON to format into a log
type Metric struct {
	// Key specifying the value to find in the json
	Key string
	// Colour specifying the colour to output the value
	Colour color.Attribute
	//
	counter map[string]int
}

// CreateMetric given a spec <colour>|<fieldName> or just <fieldName>
func CreateMetric(key, colour string) *Metric {
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

	return &Metric{
		Key:     key,
		Colour:  colourTable[colour],
		counter: make(map[string]int),
	}
}

func (m *Metric) Count(value interface{}) {
	str := fmt.Sprintf("%v", value)
	_, ok := m.counter[str]
	if !ok {
		m.counter[str] = 1
	}
	m.counter[str]++
}

// Print the value as per Field specification
func (m *Metric) Print(value interface{}) {
	var keys []string
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	// Print metric according to it's own colour 'm.Colour'
	paint := color.New(m.Colour).SprintFunc()

	for k := range m.counter {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Printf("%s:  ", blue(m.Key))
	for _, k := range keys {
		fmt.Printf("%s: %s", paint(k), red(m.counter[k]))
		fmt.Print(`  `)
	}
	fmt.Print("\n")
}
