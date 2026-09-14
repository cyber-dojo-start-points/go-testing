package hiker

import "testing"

func Test_answer_is_even(t *testing.T) {
	if have, want := answer()%2, 0; have != want {
		t.Errorf("answer modulo 2: have %d; want: %d", have, want)
	}
}

func Test_answer_is_two_digits(t *testing.T) {
	if have := answer(); have < 10 || have > 99 {
		t.Errorf("answer: have %d; want: two digits", have)
	}
}

func Test_life_the_universe_and_everything(t *testing.T) {
	if have, want := answer(), 42; have != want {
		t.Errorf("answer: have %d; want: %d", have, want)
	}
}
