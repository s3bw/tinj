package tinj

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-bongo/go-dotaccess"
)

func AggregateStdin(args, sep string) {
	metrics := ConstructMetrics(args)
	logAggregator := CreateLogAggregator(metrics, sep)

	stdin := bufio.NewReader(os.Stdin)
	for {
		nextLine, _, err := stdin.ReadLine()
		if err != nil {
			break
		}
		logAggregator.Print(nextLine)
		nextLine = nil
	}
	fmt.Print("\n")
}

// ConstructMetrics parses output format from string
func ConstructMetrics(args string) []*Metric {
	var metrics []*Metric

	// r, _ := regexp.Compile(FieldExpression)
	// for _, metricInfo:= range r.FindAllString(args, -1) {
	//	metricKey, colour := SplitFieldInfo(metricInfo)
	metric := CreateMetric(`httpRequest.status`, `red`)
	metrics = append(metrics, metric)
	//}
	return metrics
}

// LogAggregator turns a JSON input into metrics
type LogAggregator struct {
	Metrics   []*Metric
	Separator string
}

func CreateLogAggregator(metrics []*Metric, separator string) *LogAggregator {
	return &LogAggregator{Metrics: metrics, Separator: separator}
}

func (l *LogAggregator) Print(line []byte) {
	var dict map[string]interface{}

	err := json.Unmarshal(line, &dict)

	if err != nil {
		fmt.Println(string(line))
		return
	}

	for _, metric := range l.Metrics {
		value, _ := dotaccess.Get(dict, metric.Key)
		if value != nil {
			metric.Count(value)
			metric.Print(value)
			fmt.Print(l.Separator)
		}
	}
}
