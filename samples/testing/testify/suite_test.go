package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Interface to be tested
type Doubler interface {
	Double(int) int
}

// Implementation 1 (using addition)
type ByAdd struct{}

func (b ByAdd) Double(x int) int {
	return x + x
}

// Implementation 2 (using multiplication)
type ByMul struct{}

func (b ByMul) Double(x int) int {
	return x * 2
}

// Suite definition
type DoublerSuite struct {
	suite.Suite
	Doubler
}

// This is the function run by go test
func TestDoubler(t *testing.T) {
	// Create the suites with different implementations
	byAddSuite := &DoublerSuite{Doubler: ByAdd{}}
	byMulSuite := &DoublerSuite{Doubler: ByMul{}}
	// Run the suite tests on every implementation of the interface
	suite.Run(t, byAddSuite)
	suite.Run(t, byMulSuite)
}

// Test function run by a suite
func (d *DoublerSuite) TestDouble() {
	testCases := map[string]struct {
		input, expected int
	}{
		"input 1": {1, 2},
		"input 2": {2, 4},
		"input 3": {3, 6},
	}

	for name, test := range testCases {
		d.T().Run(name, func(t *testing.T) {
			result := d.Double(test.input)
			d.Equalf(test.expected, result, "double %d should be %d, got %d", test.input, test.expected, result)
		})
	}
}
