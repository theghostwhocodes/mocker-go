package filters

import (
	"testing"
)

func TestCheckArrayEquality(t *testing.T) {
	slice1 := []string{"string1", "string2", "string3"}
	slice2 := []string{"string1", "string2", "string3"}

	equals := checkArrayEquality(slice1, slice2)

	if !equals {
		t.Fail()
	}
}

func TestCheckArrayInequality1(t *testing.T) {
	slice1 := []string{"string1", "string2", "string3"}
	slice2 := []string{"string1", "string2", "string4"}

	equals := checkArrayEquality(slice1, slice2)

	if equals {
		t.Fail()
	}
}

func TestCheckArrayInequality2(t *testing.T) {
	slice1 := []string{"string1", "string2"}
	slice2 := []string{"string1", "string2", "string4"}

	equals := checkArrayEquality(slice1, slice2)

	if equals {
		t.Fail()
	}
}
