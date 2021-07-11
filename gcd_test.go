package utils

import "testing"

func TestGcdBasic(t *testing.T) {
	actual, err := Gcd(18, 69)
	if err != nil {
		t.Errorf("GcdBasic failed.\nError: %v", err)
	}

	const expected = 3

	if actual != expected {
		t.Errorf("GcdBasic failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGcdNeg(t *testing.T) {
	actual, err := Gcd(18, -69)
	if err != nil {
		t.Errorf("GcdNeg failed.\nError: %v", err)
	}

	const expected = 3

	if actual != expected {
		t.Errorf("GcdNeg failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGcdZero(t *testing.T) {
	_, err := Gcd(0, 18)
	if err == nil {
		t.Errorf("Expected GcdZero to fail.")
	}
}

func TestGcdBig(t *testing.T) {
	actual, err := Gcd(49, 9223372036854775807)
	if err != nil {
		t.Errorf("GcdBig failed.\nError: %v", err)
	}

	actual2, err2 := Gcd(49, 9223372036854775807)
	if err2 != nil {
		t.Errorf("GcdBig failed.\nError: %v", err)
	}
	if actual != actual2 {
		t.Errorf("GcdBig failed, depends on parameters order.\nValues should be identical.\n"+
			"Left: %v\nRight: %v", actual, actual2)
	}

	const expected = 49

	if actual != expected {
		t.Errorf("GcdBig failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}
