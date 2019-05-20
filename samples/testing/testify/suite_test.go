package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestDoubler(t *testing.T) {
	// Create the suites with different implementations
	byAddSuite := &DoublerSuite{
		Doubler: ByAdd{},
	}
	byMulSuite := &DoublerSuite{
		Doubler: ByMul{},
	}
	// Run the suite tests on every implementation of the interface
	suite.Run(t, byAddSuite)
	suite.Run(t, byMulSuite)
}

func (d *DoublerSuite) TestDouble() {
	r := assert.New(d.T())

	testCases := []struct {
		input, expected int
	}{
		{1, 2},
		{2, 4},
		{3, 6},
	}

	for _, test := range testCases {
		result := d.Double(test.input)
		r.Equalf(test.expected, result, "double %d should be %d, got %d", test.input, test.expected, result)
	}
}
