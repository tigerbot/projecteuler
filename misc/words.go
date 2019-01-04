package misc

import (
	"fmt"
	"math"
	"strings"
)

var (
	ones = map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
	tens = map[int]string{
		2: "twenty",
		3: "thirty",
		4: "forty",
		5: "fifty",
		6: "sixty",
		7: "seventy",
		8: "eighty",
		9: "ninety",
	}
	teens = map[int]string{
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}
	mags = map[int]string{
		3:  "thousand",
		6:  "million",
		9:  "billion",
		12: "trillion",
		15: "quadrillion",
	}
)

// nameGroup takes a number < 1000 and returns the english name.
func nameGroup(num int) string {
	if num > 999 || num < 0 {
		panic(fmt.Errorf("`nameGroup` cannot create name for %d", num))
	}
	var result string

	hundred, remainder := num/100, num%100
	if hundred != 0 {
		result += ones[hundred] + " hundred"
	}
	if hundred != 0 && remainder != 0 {
		result += " and "
	}
	if remainder != 0 {
		if teens[remainder] != "" {
			result += teens[remainder]
		} else {
			ten, one := remainder/10, remainder%10
			if ten != 0 {
				result += tens[ten]
			}
			if ten != 0 && one != 0 {
				result += "-"
			}
			if one != 0 {
				result += ones[one]
			}
		}
	}

	return result
}

// NameNumber returns the English name for the provided number.
func NameNumber(num int) string {
	var result string
	if num == 0 {
		return "zero"
	}
	if num < 0 {
		result = "negative "
		num *= -1
	}

	first := true
	for mag := 3 * int(math.Floor(math.Log10(float64(num))/3)); mag >= 0; mag -= 3 {
		pow := int(math.Pow10(int(mag)))
		if grpName := nameGroup(num / pow); grpName != "" {
			if !first {
				result += ", "
			}
			result += fmt.Sprintf("%s %s", grpName, mags[mag])
			first = false
		}
		num = num % pow
	}
	return result
}

// ScoreWord adds the "score" of each character in a name
func ScoreWord(name string) int {
	buf := []byte(strings.ToUpper(name))

	var sum int
	for _, b := range buf {
		sum += int(b-'A') + 1
	}

	return sum
}
