package primes

import "sort"

type sortableInt64 []int64

func (s sortableInt64) Len() int           { return len(s) }
func (s sortableInt64) Less(i, j int) bool { return s[i] < s[j] }
func (s sortableInt64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func sortInt64(list []int64) {
	sort.Sort(sortableInt64(list))
}

// Pow is an integer version of the math.Pow function. It utilizes exponentiation by squaring.
func Pow(a, b int64) int64 {
	p := int64(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
