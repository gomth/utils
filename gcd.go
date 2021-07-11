package utils

import "errors"

// Gcd - recursive "greatest common divisor" Euclidean algo impl
// Beware of Fibonacci numbers.
func Gcd(a, b int64) (int64, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("Unexpected zero input")
	}

	if a > b {
		return gcd(a, b), nil
	}

	return gcd(b, a), nil
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
