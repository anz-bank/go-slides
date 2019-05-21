package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	// You can create many test cases succinctly with an anonymous struct array
	testCases := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: "a b c", sep: " ", want: []string{"a", "b", "c"}},
		"Wrong sep":    {input: "a/b/c", sep: " ", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: " ", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c", ""}},
	}

	for name, test := range testCases {
		// t.Run creates a sub test and runs it like a normal test
		t.Run(name, func(t *testing.T) {
			result := strings.Split(test.input, test.sep)
			if !reflect.DeepEqual(result, test.want) {
				t.Errorf("expected %v, got %v", test.want, result)
			}
		})
	}
}
