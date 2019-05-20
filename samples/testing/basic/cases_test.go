package basic

import "testing"

func TestCases(t *testing.T) {
	// You can create many test cases succinctly with an anonymous struct array
	testCases := []struct {
		input, expected string
	}{
		{"Hello", "Hello M8"},
		{"Goodbye", "Goodbye M8"},
		{"💥💥💥", "💥💥💥 M8"},
		{"", " M8"},
	}

	for _, test := range testCases {
		result := AddM8(test.input)
		if result != test.expected {
			// If one fails, log the failure and fail the test
			t.Logf("wanted %s, got %s", test.expected, result)
			t.Fail()
		}
	}
}

func AddM8(str string) string {
	return str + " M8"
}
