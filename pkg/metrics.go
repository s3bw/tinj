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
	v := value.(float64)
	s := fmt.Sprintf("%.0f", v)
	_, ok := m.counter[s]
	if !ok {
		m.counter[s] = 1
	}
	m.counter[s]++
}

// Print the value as per Field specification
func (m *Metric) Print(value interface{}) {
	color.Set(m.Colour)
	// fmt.Print(m.Key)
	var keys []string

	for k := range m.counter {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Print("\r")
	for _, k := range keys {
		fmt.Printf("%s: count %d", k, m.counter[k])
		color.Set(color.FgCyan)
		fmt.Print(` | `)
		color.Set(m.Colour)
	}
	color.Unset()
}
