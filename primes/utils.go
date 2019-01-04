package primes

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
