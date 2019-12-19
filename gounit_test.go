package gounit

import (
	"testing"
)

func TestAssertEquals(t *testing.T) {

	AssertEquals(t, 1, 1)
	AssertEquals(t, "s1", "s1")
	AssertEquals(t, 1.03, 1.03)

}

func TestAssertNotEquals(t *testing.T) {

	AssertNotEquals(t, 1, 2)
	AssertNotEquals(t, "s1", "s2")
	AssertNotEquals(t, 1.04, 1.03)

}

func TestAssertTrue(t *testing.T) {

	condition1 := 1 < 2
	AssertTrue(t, "Falhou", condition1)

}

func TestAssertFalse(t *testing.T) {

	condition1 := 2 < 1
	AssertFalse(t, "Falhou", condition1)

}
