package primes

import (
	"reflect"
	"testing"
)

func TestBetween(t *testing.T) {
	if list := Between(1, 10); !reflect.DeepEqual(list, []int{2, 3, 5, 7}) {
		t.Errorf("expected primes from 1-10 to be [2,3,5,7]; got %d", list)
	}
	if list := Between(2, 7); !reflect.DeepEqual(list, []int{2, 3, 5, 7}) {
		t.Errorf("expected primes from 2-7 to be [2,3,5,7]; got %d", list)
	}
	if list := Between(7, 8); !reflect.DeepEqual(list, []int{7}) {
		t.Errorf("expected primes from 7-8 to be [7]; got %d", list)
	}
	if list := Between(7, 7); !reflect.DeepEqual(list, []int{7}) {
		t.Errorf("expected primes from 7-7 to be [7]; got %d", list)
	}

}

func TestExpand(t *testing.T) {
	resetSieve()
	if list := Between(2, 255); list[len(list)-1] != 251 {
		t.Errorf("expected primes from 2-255 to end in 7069; got %d", list[len(list)-1])
	}
	if list := Between(2, 255); list[len(list)-1] != 251 {
		t.Errorf("expected primes from 2-255 to end in 251; got %d", list[len(list)-1])
	}

	// Not interested in the results so much as making sure it doesn't crash
	if !testing.Short() {
		Between(1, 1<<26)
	}
}
