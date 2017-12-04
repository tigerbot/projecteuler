package primes

import (
	"fmt"
	"math"
)

// Factor returns a slice of all of the prime factors for the given number.
func Factor(num int64) []int64 {
	if num == 1 {
		return []int64{1}
	}
	if num < 1 {
		panic(fmt.Errorf("cannot factor non-natural number %d", num))
	}

	result := []int64{}
	loopPrimes := func(list []int64) {
		for _, prime := range list {
			if prime > num {
				panic(fmt.Sprintf("factoring %d reached prime %d, but only %d remaining", num, prime, num))
			}
			for num%prime == 0 {
				result = append(result, prime)
				num /= prime
			}
			// We've already removed all of the smaller prime factors, so if the number is bigger
			// than the square of the current prime no larger primes can be factors
			if prime*prime > num {
				result = append(result, num)
				num = 1
			}
			if num == 1 {
				return
			}
		}
	}

	// Pull out all the factors of the primes we've already discovered before determining what the
	// largest prime factor can be in case we can make that smaller.
	loopPrimes(cachedSieve)
	if num == 1 {
		return result
	}

	// Then loop through all of the primes between what we had cached and the square root of the
	// number we haven't been able to reduce yet.
	loopPrimes(Between(cachedRange, int64(math.Sqrt(float64(num))+1)))
	if num != 1 {
		result = append(result, num)
	}

	return result
}

// FactorMap returns a map of how many times each factor appears in the prime factorization of
// the given number. For example FactorMap(16) would return {2: 4} and FactorMap(60) would return
// {2: 2, 3: 1, 5: 1}
func FactorMap(num int64) map[int64]int64 {
	factors := Factor(num)
	result := map[int64]int64{}
	for _, prime := range factors {
		result[prime]++
	}
	return result
}

// CountDivisors counts the number of unique natural numbers greater than 1 that divide evenly into
// the given number without a remainder.
func CountDivisors(num int64) int64 {
	result := int64(1)
	for _, cnt := range FactorMap(num) {
		result *= (cnt + 1)
	}
	return result
}

func multiplySlice(factor int64, slice []int64) []int64 {
	result := make([]int64, len(slice))
	for i := range slice {
		result[i] = factor * slice[i]
	}
	return result
}

// Divisors finds all numbers that divide evenly into the provided number. Note that this does
// require more work than determining how many there are, so only use this function if you need the
// actual values of the divisors.
func Divisors(num int64) []int64 {
	result := []int64{1}
	for prime, cnt := range FactorMap(num) {
		prev := result
		for i, factor := int64(1), prime; i <= cnt; i++ {
			result = append(result, multiplySlice(factor, prev)...)
			factor *= prime
		}
	}
	sortInt64(result)
	return result
}

// SumDivisors adds all of the numbers less than the provided that divide evenly into it.
func SumDivisors(num int64) int64 {
	divs := Divisors(num)
	divs = divs[:len(divs)-1]

	var sum int64
	for _, val := range divs {
		sum += val
	}
	return sum
}

// EulerPhi calculates the number of positive integers less than n that are relatively prime to n.
func EulerPhi(num int64) int64 {
	result := int64(1)
	for prime, cnt := range FactorMap(num) {
		result *= (prime - 1) * Pow(prime, cnt-1)
	}
	return result
}

// LCM returns the Least Common Multiple for all of the numbers specified
func LCM(values ...int64) int64 {
	allFactors := map[int64]int64{}
	for _, val := range values {
		for prime, count := range FactorMap(val) {
			if allFactors[prime] < count {
				allFactors[prime] = count
			}
		}
	}

	result := int64(1)
	for prime, count := range allFactors {
		result *= Pow(prime, count)
	}
	return result
}

// GCD return the Greatest Common Divisor for all of the numbers specified
func GCD(values ...int64) int64 {
	commonFactors := FactorMap(values[0])
	for _, val := range values {
		newFactors := FactorMap(val)
		for prime := range commonFactors {
			if newFactors[prime] < commonFactors[prime] {
				commonFactors[prime] = newFactors[prime]
			}
			if commonFactors[prime] == 0 {
				delete(commonFactors, prime)
			}
		}
	}

	result := int64(1)
	for prime, count := range commonFactors {
		result *= Pow(prime, count)
	}
	return result
}

func combineFactors(factors []int64, count int) []int64 {
	if count == 0 || len(factors) < count {
		return nil
	}
	if count == 1 {
		return factors
	}

	var result []int64
	for ind, prime := range factors {
		subRes := combineFactors(append([]int64(nil), factors[:ind]...), count-1)
		for i := range subRes {
			subRes[i] *= prime
		}
		result = append(result, subRes...)
	}
	return result
}
func countDivisbleBy(limit int64, factors []int64) int64 {
	var result int64
	for i := 1; i <= len(factors); i++ {
		var subRes int64
		for _, num := range combineFactors(factors, i) {
			subRes += (limit - 1) / num
		}

		if i%2 == 1 {
			result += subRes
		} else {
			result -= subRes
		}
	}
	return result
}
