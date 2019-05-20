package basic

import (
	"fmt"
	"testing"
)

// Test functions are prefixed Test and take *testing.T as input
func TestLog(t *testing.T) {
	t.Logf("log test output with t.Log and t.Logf")
	// This test will pass
}

func TestFail(t *testing.T) {
	// Mark a test as failed with t.Fail()
	t.Fail()

	// This test will fail, but continue to execute
	t.Log("Test continued execution after failing")
}

func TestFatal(t *testing.T) {
	// Log, Fail and end a test early with t.Fatal()
	t.Fatalf("This test failed and cannot proceed")

	// No further execution will occur
	t.Log("THIS SHOULD NOT BE PRINTED")
}

func TestExample(t *testing.T) {
	val, err := functionToTest(-1)
	if err != nil {
		t.Fatal("Error should not have occured")
	}

	expected := 1
	if val != expected {
		t.Log("Func returned unexpected value")
		t.Fail()
	}
}

func functionToTest(i int) (int, error) {
	if i > 0 {
		return i, nil
	}
	return 0, fmt.Errorf("Negative :(")
}
