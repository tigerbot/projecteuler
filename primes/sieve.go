package primes

import (
	"math"
	"sort"
)

var (
	cachedSieve = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	cachedRange = int(30)
)

// resetSieve is for testing purposes only.
func resetSieve() {
	cachedSieve = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	cachedRange = int(30)
}

func expandSieve(num int) {
	if num < cachedRange {
		return
	}
	const minDiff, maxDiff = int(200), int(1 << 23)
	for num-cachedRange > maxDiff {
		expandSieve(cachedRange + maxDiff)
	}
	if num-cachedRange < minDiff {
		num = cachedRange + minDiff
	}

	// Establish all of the prime candidates in the range we haven't checked yet
	var start int
	if cachedRange%2 == 0 {
		start = cachedRange + 1
	} else {
		start = cachedRange
	}
	candidates := make(map[int]bool, (num-cachedRange)/2+1)
	for ind := start; ind <= num; ind += 2 {
		candidates[ind] = true
	}

	// First go through our existing primes and remove all multiples of those
	for _, prime := range cachedSieve {
		for ind := ((cachedRange-1)/prime + 1) * prime; ind <= num; ind += prime {
			candidates[ind] = false
		}
	}

	// Then go through all new number in ascending order and store everything that's still a candidate
	for ind := start; ind <= num; ind += 2 {
		if !candidates[ind] {
			continue
		}

		cachedSieve = append(cachedSieve, ind)
		for j := 2 * ind; j <= num; j += ind {
			candidates[j] = false
		}
	}
	cachedRange = num
}

// IsPrime checks to see if the specified number is prime.
func IsPrime(num int) bool {
	if cachedRange >= num {
		l := len(cachedSieve)
		ind := sort.Search(l, func(i int) bool { return cachedSieve[i] >= num })
		return ind < l && cachedSieve[ind] == num
	}

	limit := int(math.Sqrt(float64(num)))
	expandSieve(limit)
	for _, prime := range cachedSieve {
		if num%prime == 0 {
			return false
		}
		if prime > limit {
			break
		}
	}
	return true
}

// PrimeIndex takes a number, and if it is prime it will return its index. It is
// the inverse of NthPrime
func PrimeIndex(num int) int {
	if num < cachedRange {
		expandSieve(num)
	}
	l := len(cachedSieve)
	ind := sort.Search(l, func(i int) bool { return cachedSieve[i] >= num })
	if ind >= l || cachedSieve[ind] != num {
		return -1
	}
	// account for 0-indexing of the array
	return ind + 1
}

// NthPrime returns the value of the nth prime (starting with 2 as the 1st prime).
func NthPrime(num int) int {
	// acount for 0-indexing of the array
	num--
	for rng := cachedRange + 1e3; num >= len(cachedSieve); rng += 1e3 {
		expandSieve(rng)
	}
	return cachedSieve[num]
}

// Between returns a list of all primes in the specified range
func Between(lower, upper int) []int {
	if lower > upper {
		return nil
	}
	expandSieve(upper)
	ind1 := sort.Search(len(cachedSieve), func(i int) bool { return cachedSieve[i] >= lower })
	if ind1 < 0 {
		ind1 = 0
	}
	ind2 := sort.Search(len(cachedSieve), func(i int) bool { return cachedSieve[i] >= upper })
	if ind2 > len(cachedSieve) {
		ind2 = len(cachedSieve)
	} else if ind2 < len(cachedSieve)-1 && cachedSieve[ind2] == upper {
		ind2++
	}
	return append([]int(nil), cachedSieve[ind1:ind2]...)
}
