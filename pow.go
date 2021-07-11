package utils

// Modular multiplication
func mulm(a, b, m int64) int64 {
	if b == 1 {
		return a
	}

	if b%2 == 0 {
		t := mulm(a, b/2, m)
		return (2 * t) % m
	}
	return (mulm(a, b-1, m) + a) % m
}

// Powm - Modular exponentiation
func Powm(a, b, m int64) int64 {
	if b == 0 {
		return 1
	}
	if b%2 == 0 {
		t := Powm(a, b/2, m)
		return mulm(t, t, m) % m
	}
	return (mulm(Powm(a, b-1, m), a, m)) % m
}
