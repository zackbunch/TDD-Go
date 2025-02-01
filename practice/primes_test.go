package practice

import "testing"

func TestPrimes(t *testing.T) {
	got := IsPrime(49)
	want := false

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
