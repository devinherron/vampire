package txtreader

import "testing"

func Test_GetPrompt(t *testing.T) {
	testStrings := []struct {
		testString     string
		expectedPrompt int
		expectedEntry  int
	}{
		{"No match", 0, 0},
		{"23a", 23, 0},
		{"17", 0, 0},
		{"19d", 0, 0},
		{"1c", 1, 2},
	}

	for _, test := range testStrings {
		prompt, entry := GetPrompt(test.testString)
		if prompt != test.expectedPrompt || entry != test.expectedEntry {
			t.Errorf("GetPrompt(%s) was incorrect, got prompt: %d, entry: %d, expected prompt: %d, entry: %d",
				test.testString, prompt, entry, test.expectedPrompt, test.expectedEntry)
		}
	}
}
