package tinj

import (
	"testing"

	"github.com/fatih/color"
)

var deconstructFormatTest = []struct {
	// The format provided
	givenFormat string
	// The fields expected from the format
	expectedFields []*Field
}{
	// Case given field and colour
	{"(message|red)", []*Field{{Key: "message", Colour: color.FgRed}}},
	// Case given just a field
	{"(status)", []*Field{{Key: "status", Colour: color.FgWhite}}},
	// Case given a nested field
	{"(status.field)", []*Field{{Key: "status.field", Colour: color.FgWhite}}},
	// Case given a nested field and a colour
	{"(status.field|black)", []*Field{{Key: "status.field", Colour: color.FgBlack}}},
	// Case given multiple fields
	{"(status),(package|yellow)", []*Field{
		{Key: "status", Colour: color.FgWhite},
		{Key: "package", Colour: color.FgYellow},
	}},
}

func TestDeconstructFormat(t *testing.T) {
	for _, tt := range deconstructFormatTest {
		result := ConstructFields(tt.givenFormat)

		if len(result) != len(tt.expectedFields) {
			t.Errorf("Expected: %v, got: %v", tt.expectedFields, result)
		}

		for i := range result {
			if result[i].Key != tt.expectedFields[i].Key || result[i].Colour != tt.expectedFields[i].Colour {
				t.Errorf("Expected: %v, got: %v", tt.expectedFields[i], result[i])
			}
		}
	}
}

var splitIntoFieldsTest = []struct {
	fieldInfo      string
	expectedField  string
	expectedColour string
}{
	// Case given a field
	{"(message)", "message", DefaultColour},
	// Case given a field and a colour
	{"(message|red)", "message", "red"},
	// Case given a nested field and a colour
	{"(nested.field|red)", "nested.field", "red"},
}

func TestSplitFieldInfo(t *testing.T) {
	for _, tt := range splitIntoFieldsTest {
		resultField, resultColour := SplitFieldInfo(tt.fieldInfo)

		if resultField != tt.expectedField {
			t.Errorf("Expected: %s, got: %s", tt.expectedField, resultField)
		}

		if resultColour != tt.expectedColour {
			t.Errorf("Expected: %s, got: %s", tt.expectedColour, resultColour)
		}
	}
}
