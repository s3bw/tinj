package tinj

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/go-bongo/go-dotaccess"
)

func AggregateStdin(args, sep string) {
	metrics := ConstructMetrics(args)
	lineAggregator := CreateLineAggregator(metrics, sep)

	stdin := bufio.NewReader(os.Stdin)
	for {
		nextLine, _, err := stdin.ReadLine()
		if err != nil {
			break
		}
		lineAggregator.Print(nextLine)
		nextLine = nil
	}
	fmt.Print("\n\n")
}

// ConstructMetrics parses output format from string
func ConstructMetrics(args string) []*Metric {
	var metrics []*Metric

	r, _ := regexp.Compile(FieldExpression)
	for _, metricInfo := range r.FindAllString(args, -1) {
		metricKey, colour := SplitFieldInfo(metricInfo)
		metric := CreateMetric(metricKey, colour)
		metrics = append(metrics, metric)
	}
	return metrics
}

// LogAggregator turns a JSON input into metrics
type LineAggregator struct {
	Metrics   []*Metric
	Separator string
}

func CreateLineAggregator(metrics []*Metric, separator string) *LineAggregator {
	return &LineAggregator{Metrics: metrics, Separator: separator}
}

func (l *LineAggregator) Print(line []byte) {
	var dict map[string]interface{}

	err := json.Unmarshal(line, &dict)

	if err != nil {
		fmt.Println(string(line))
		return
	}

	removeLine := 0
	for _, metric := range l.Metrics {
		value, _ := dotaccess.Get(dict, metric.Key)
		if value != nil {
			metric.Count(value)
			metric.Print(value)
			fmt.Print("\n")

			removeLine++
		}
	}
	// Magic charater to move to the start of line above
	for i := 1; i <= removeLine; i++ {
		fmt.Print("\033[F")
	}
}
