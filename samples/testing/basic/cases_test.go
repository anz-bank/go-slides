package main

// Attributed to Dave Cheney
// https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

import (
	"testing"
)

func IsEven(n int) bool {
	return n%2 == 0
}

func TestIsEven(t *testing.T) {
	// You can create many test cases succinctly with an anonymous struct array
	testCases := map[string]struct {
		input int
		want  bool
	}{
		"even":     {input: 132, want: true},
		"odd":      {input: 3, want: false},
		"0":        {input: 0, want: true},
		"negative": {input: -43, want: false},
	}

	for name, test := range testCases {
		// t.Run creates a sub test and runs it like a normal test
		t.Run(name, func(t *testing.T) {
			result := IsEven(test.input)
			if result != test.want {
				t.Errorf("expected %v, got %v", test.want, result)
			}
		})
	}
}
