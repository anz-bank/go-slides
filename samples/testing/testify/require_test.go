package testify

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequire(t *testing.T) {
	// require.New creates an object that can be used lik more common test frameworks
	r := require.New(t)
	r.Equal(1, 1, "1 should be equal to 1")
}

func TestRequireContains(t *testing.T) {
	r := require.New(t)

	// Contain can by used on strings
	r.Contains("Hello world", "Hello", "'Hello world' should contain 'Hello'")

	// slices/array
	r.Contains([]int{1, 2}, 1, "Array should contain 1")

	// and maps
	r.Contains(map[int]int{1: 2}, 1, "Map should contain 1 as a key")
}

func TestRequireErr(t *testing.T) {
	r := require.New(t)

	// NoError checks that a returned error is nil
	var err error
	r.NoError(err, "Nil error should pass NoError()")

	// Error checks that an error is not nil
	err = fmt.Errorf("Message")
	r.Error(err, "non-nil error should pass Error")

	// You still need to check the message separately
	r.Contains(err.Error(), "Message", "Error message should contain 'Message'")

	// Or you can use an error equal that implicitly checks the error message
	r.EqualError(err, "Message", "ErrorEqual should check the messages")
}

func TestRequireEqual(t *testing.T) {
	r := require.New(t)

	// Equal checks the values referenced by pointers
	x := 1
	y := 1
	r.Equal(&x, &y, "Equal is too smart")

	// EqualValues can be used when similar types are involved
	r.EqualValues(int32(1), uint64(1), "EqualValues is definitely too smart")
}
