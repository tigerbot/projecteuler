package misc

import (
	"fmt"
	"math"
)

// FindCycle finds the repeating part of the division of the two specified integers.
func FindCycle(num, den int) []int {
	if num < 1 || den < 1 {
		panic(fmt.Errorf("must provide position numbers to FindCycle"))
	}
	remainders := make(map[int]int)
	decimals := make([]int, 0)

	left := num
	for ind := 0; ind < 1e5 && left > 0; ind++ {
		for left < den {
			decimals = append(decimals, 0)
			left *= 10
		}
		if remainders[left] != 0 {
			break
		}
		remainders[left] = len(decimals)
		decimals = append(decimals, left/den)
		left = (left % den) * 10
	}

	if left == 0 {
		return nil
	}
	if remainders[left] == 0 {
		panic(fmt.Errorf("failed to find cycle: %v", decimals))
	}
	return decimals[remainders[left]:]
}

// SplitDigits returns a list of the digits used to represent the number in decimal notation. The
// first element in the array represents the highest magnitude.
func SplitDigits(num, base int) []int {
	if num < 0 || base < 0 {
		panic("Tried to split negative number, overflow suspected")
	}
	if num == 0 {
		return []int{0}
	}
	cnt := int(math.Floor(math.Log10(float64(num))/math.Log10(float64(base))) + 1)
	result := make([]int, cnt)
	for i := 1; i <= cnt; i++ {
		result[cnt-i] = int(num % base)
		num /= base
	}
	return result
}

// MergeDigits is the inverse of the SplitDigit function.
func MergeDigits(digits []int, base int) int {
	var result int
	for _, dig := range digits {
		result = base*result + int(dig)
	}
	return result
}

func IsPandigit(digits []int) bool {
	used := make(map[int]bool, len(digits))
	for _, val := range digits {
		if used[val] || val == 0 || val > len(digits) {
			return false
		}
		used[val] = true
	}
	return true
}
