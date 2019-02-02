package primes

import "math"

// Pow is an integer version of the math.Pow function. It utilizes exponentiation by squaring.
func Pow(a, b int) int {
	p := int(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// PowMod is similar to Pow, but does modular exponentation. It returns (a^b)%m
func PowMod(a, b, m int) int {
	p := int(1)
	for b > 0 {
		if b&1 != 0 {
			p = (p * a) % m
		}

		b >>= 1
		a = (a * a) % m
	}
	return p
}

// IsSquare tests to see if an integer value is the square of another integer.
func IsSquare(x int) bool {
	if h := x & 0xf; h != 0 && h != 1 && h != 4 && h != 9 {
		return false
	}
	sqr := int(math.Sqrt(float64(x)))
	return sqr*sqr == x
}
