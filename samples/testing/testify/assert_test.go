package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssert(t *testing.T) {
	// require.New creates an object that can be used like more common test frameworks
	a := assert.New(t)
	a.Equal(1, 1, "1 should be equal to 1")
}

func TestAssertVsRequire(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	// Assert and require work very similary
	a.Equal(1, 2, "Not Equal")

	// Require methods stop execution if they fail
	r.Equal(1, 2, "Not equal")

	t.Log("THIS WILL NOT PRINT")
}

func TestAssertContains(t *testing.T) {
	a := assert.New(t)

	// Contain can by used on strings
	a.Contains("Hello world", "Hello", "'Hello world' should contain 'Hello'")

	// slices/array
	a.Contains([]int{1, 2}, 1, "Array should contain 1")

	// and maps
	a.Contains(map[int]int{1: 2}, 1, "Map should contain 1 as a key")
}

func TestAssertErr(t *testing.T) {
	a := assert.New(t)

	// NoError checks that a returned error is nil
	var err error
	a.NoError(err, "Nil error should pass NoError()")

	// Error checks that an error is not nil
	err = fmt.Errorf("Message")
	a.Error(err, "non-nil error should pass Error")

	// You still need to check the message separately
	a.Contains(err.Error(), "Message", "Error message should contain 'Message'")

	// Or you can use an error equal that implicitly checks the error message
	a.EqualError(err, "Message", "ErrorEqual should check the messages")
}

func TestAssertJSONEq(t *testing.T) {
	a := assert.New(t)
	JSON1 := `{"key": "val", "foo": "bar"}`
	JSON2 := `{"foo": "bar", "key": "val"}`

	// JSONEq is very convenient for checking json as raw strings
	// Does not care about ordering
	a.JSONEq(JSON1, JSON2, "JSONEq failed")
}

func TestAssertEqual(t *testing.T) {
	a := assert.New(t)

	// Equal checks the values referenced by pointers
	x := 1
	y := 1
	a.Equal(&x, &y, "Equal is too smart")

	// EqualValues can be used when similar types are involved
	a.EqualValues(int32(1), uint64(1), "EqualValues is definitely too smart")

	// assert methods also return a boolean you can use to terminate if needed
	// For this reason, require is not used very often
	err := dangerousIfError()
	if !a.NoError(err, "Error executing crucial function") {
		a.FailNow("Test cannot proceed")
	}
}

func dangerousIfError() error {
	return fmt.Errorf("Bad things")
}
