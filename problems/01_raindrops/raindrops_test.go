package raindrops

import "testing"

var testCases = []struct {
	description string
	input       int
	expected    string
}{
	{"3 only", 3, "Pling"},
	{"5 only", 5, "Plang"},
	{"7 only", 7, "Plong"},
	{"3 and 5, not 7", 15, "PlingPlang"},
	{"3 and 7, not 5", 21, "PlingPlong"},
	{"5 and 7, not 3", 35, "PlangPlong"},
	{"3, 5, and 7", 105, "PlingPlangPlong"},
	{"none, plain number", 34, "34"},
}

func TestConvert(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Convert(tc.input); actual != tc.expected {
				t.Fatalf("Convert(%d) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}
