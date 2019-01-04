package primes

import (
	"fmt"
	"math"
	"sort"
)

// Factor returns a slice of all of the prime factors for the given number.
func Factor(orig int) []int {
	if orig == 1 {
		return []int{1}
	}
	if orig < 1 {
		panic(fmt.Errorf("cannot factor non-natural number %d", orig))
	}
	num := orig

	result := []int{}
	loopPrimes := func(list []int) {
		for _, prime := range list {
			if prime > num {
				panic(fmt.Sprintf("factoring %d reached prime %d, but only %d remaining", orig, prime, num))
			}
			for num%prime == 0 {
				result = append(result, prime)
				num /= prime
			}
			// We've already removed all of the smaller prime factors, so if the number is bigger
			// than the square of the current prime what's left must be a prime.
			if num != 1 && prime*prime > num {
				result = append(result, num)
				num = 1
			}
			if num == 1 {
				return
			}
		}
	}

	// Pull out all the factors of the primes we've already discovered before trying to expand the
	// sieve. This helps limit how big we have to make the sieve when dealing with larger numbers.
	loopPrimes(cachedSieve)
	if num == 1 {
		return result
	}

	// Then loop through all of the primes between what we had cached and the square root of the
	// number we haven't been able to reduce yet.
	loopPrimes(Between(cachedRange, int(math.Sqrt(float64(num))+1)))
	if num != 1 {
		result = append(result, num)
	}

	return result
}

// FactorMap returns a map of how many times each factor appears in the prime factorization of
// the given number. For example FactorMap(16) would return {2: 4} and FactorMap(60) would return
// {2: 2, 3: 1, 5: 1}
func FactorMap(num int) map[int]int {
	factors := Factor(num)
	result := map[int]int{}
	for _, prime := range factors {
		result[prime]++
	}
	return result
}

// CountDivisors counts the number of unique natural numbers greater than 1 that divide evenly into
// the given number without a remainder.
func CountDivisors(num int) int {
	result := int(1)
	for _, cnt := range FactorMap(num) {
		result *= (cnt + 1)
	}
	return result
}

func multiplySlice(factor int, slice []int) []int {
	result := make([]int, len(slice))
	for i := range slice {
		result[i] = factor * slice[i]
	}
	return result
}

// Divisors finds all numbers that divide evenly into the provided number. Note that this does
// require more work than determining how many there are, so only use this function if you need the
// actual values of the divisors.
func Divisors(num int) []int {
	result := []int{1}
	for prime, cnt := range FactorMap(num) {
		prev := result
		for i, factor := int(1), prime; i <= cnt; i++ {
			result = append(result, multiplySlice(factor, prev)...)
			factor *= prime
		}
	}
	sort.Ints(result)
	return result
}

// SumDivisors adds all of the numbers less than the provided that divide evenly into it.
func SumDivisors(num int) int {
	divs := Divisors(num)
	divs = divs[:len(divs)-1]

	var sum int
	for _, val := range divs {
		sum += val
	}
	return sum
}

// EulerPhi calculates the number of positive integers less than n that are relatively prime to n.
func EulerPhi(num int) int {
	result := int(1)
	for prime, cnt := range FactorMap(num) {
		result *= (prime - 1) * Pow(prime, cnt-1)
	}
	return result
}

// LCM returns the Least Common Multiple for the numbers specified
func LCM(a, b int) int {
	d := GCD(a, b)
	return a * b / d
}

// GCD return the Greatest Common Divisor for the numbers specified
func GCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
