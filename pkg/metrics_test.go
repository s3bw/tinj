package tinj

import (
	"testing"
)

var metricValuesTest = []struct {
	// A stream of values
	values []interface{}
	// The expected count of value
	expectedCount int
	// The value that is counted
	valueCounted string
}{
	// Case given list of values
	{[]interface{}{"100", "100", "100"}, 3, "100"},
	{[]interface{}{"200", "100", "100"}, 1, "200"},
	{[]interface{}{"200", "300", "300"}, 2, "300"},
	// Test when metric is an integer
	{[]interface{}{200, 300, 300}, 2, "300"},
}

func TestCountMetrics(t *testing.T) {
	for _, tt := range metricValuesTest {
		metric := CreateMetric("statusCode", "white")

		// Count values
		for _, v := range tt.values {
			metric.Count(v)
		}

		// Check result of counted value
		result := metric.counter[tt.valueCounted]
		if result != tt.expectedCount {
			t.Errorf("Expected: %v, got: %v", tt.expectedCount, result)
		}
	}
}
