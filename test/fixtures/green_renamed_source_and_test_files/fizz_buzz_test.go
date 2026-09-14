package fizzbuzz

import "testing"

func Test_multiples_of_three_and_five_are_fizzbuzz(t *testing.T) {
	if have, want := fizzBuzz(15), "FizzBuzz"; have != want {
		t.Errorf("fizzBuzz(15): have %q; want: %q", have, want)
	}
}
