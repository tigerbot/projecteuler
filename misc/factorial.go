package misc

import (
	"fmt"
	"math/big"
)

// Factorial calculate n!, panicing if the number is too big to fit inside a 64-bit int.
func Factorial(num int64) int64 {
	var result big.Int
	result.MulRange(2, num)
	if !result.IsInt64() {
		panic(fmt.Errorf("%d! does not fit inside a 64-bit int", num))
	}
	return result.Int64()
}

// Choose calculates the number of possible combinations when choosing "choices" items from a
// pool of "total" items where order in which they are chosen does not matter
func Choose(total, choices int64) int64 {
	var result big.Int
	result.Binomial(total, choices)
	if !result.IsInt64() {
		panic(fmt.Errorf("(%d %d) does not fit inside a 64-bit int", total, choices))
	}
	return result.Int64()
}
