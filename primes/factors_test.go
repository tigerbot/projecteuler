package primes

import (
	"testing"
)

func TestGCD(t *testing.T) {
	type s struct {
		a, b, r int
	}
	expected := []s{
		{8 * 5, 8 * 54, 8},
		{27 * 7, 27 * 10, 27},
	}

	for _, e := range expected {
		if r := GCD(e.a, e.b); r != e.r {
			t.Errorf("GCD(%d, %d) returned %d, expected %d", e.a, e.b, r, e.r)
		}
		if r := GCD(e.b, e.a); r != e.r {
			t.Errorf("GCD(%d, %d) returned %d, expected %d", e.b, e.a, r, e.r)
		}
	}
}

func TestLCM(t *testing.T) {
	type s struct {
		a, b, r int
	}
	expected := []s{
		{26 * 72, 26 * 35, 26 * 72 * 35},
	}

	for _, e := range expected {
		if r := LCM(e.a, e.b); r != e.r {
			t.Errorf("LCM(%d, %d) returned %d, expected %d", e.a, e.b, r, e.r)
		}
		if r := LCM(e.b, e.a); r != e.r {
			t.Errorf("LCM(%d, %d) returned %d, expected %d", e.b, e.a, r, e.r)
		}
	}
}
