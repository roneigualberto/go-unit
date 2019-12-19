package gounit

import (
	"fmt"
	"testing"
)

// AssertEquals assertion
func AssertEquals(t *testing.T, atual, expected interface{}) {
	condition := atual == expected
	message := fmt.Sprintf("Expected value %v, but the result found was %v", expected, atual)
	AssertTrue(t, message, condition)
}

// AssertNotEquals assertion
func AssertNotEquals(t *testing.T, atual, expected interface{}) {

	condition := atual != expected
	message := fmt.Sprintf("Expected value should be different %v, but the result was equal to % v ", expected, atual)
	AssertTrue(t, message, condition)

}

//AssertTrue assertion
func AssertTrue(t *testing.T, message string, condition bool) {
	if !condition {
		Fail(t, message)
	}
}

//AssertFalse assertion
func AssertFalse(t *testing.T, message string, condition bool) {
	AssertTrue(t, message, !condition)
}

// Fail execute t.Errorf
func Fail(t *testing.T, message string, args ...interface{}) {
	t.Errorf(message, args...)
}
