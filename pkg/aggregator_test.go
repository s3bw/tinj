package tinj

import (
	"testing"

	"github.com/fatih/color"
)

var deconstructMetricFormatTest = []struct {
	// The format provided
	givenFormat string
	// The metrics expected from the format
	expectedMetrics []*Metric
}{
	// Case given metric and colour
	{"(message|red)", []*Metric{{Key: "message", Colour: color.FgRed}}},
	// Case given just a metric
	{"(status)", []*Metric{{Key: "status", Colour: color.FgWhite}}},
	// Case given a nested metric
	{"(status.field)", []*Metric{{Key: "status.field", Colour: color.FgWhite}}},
	// Case given a nested metric and a colour
	{"(status.field|black)", []*Metric{{Key: "status.field", Colour: color.FgBlack}}},
	// Case given multiple metrics
	{"(status),(package|yellow)", []*Metric{
		{Key: "status", Colour: color.FgWhite},
		{Key: "package", Colour: color.FgYellow},
	}},
}

func TestDeconstructMetricFormat(t *testing.T) {
	for _, tt := range deconstructMetricFormatTest {
		result := ConstructMetrics(tt.givenFormat)

		if len(result) != len(tt.expectedMetrics) {
			t.Errorf("Expected: %v, got: %v", tt.expectedMetrics, result)
		}

		for i := range result {
			if result[i].Key != tt.expectedMetrics[i].Key || result[i].Colour != tt.expectedMetrics[i].Colour {
				t.Errorf("Expected: %v, got: %v", tt.expectedMetrics[i], result[i])
			}
		}
	}
}
