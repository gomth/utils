package utils

import "testing"

func TestPowmBasic(t *testing.T) {
	actual := Powm(2, 3, 5)

	const expected = 3

	if actual != expected {
		t.Errorf("GcdBasic failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}
