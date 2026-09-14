package hiker

import "testing"

func Test_life_the_universe_and_everything(t *testing.T) {
	if have, want := answer(), 42; have != want {
		t.Errorf("answer: have %d; want: %d", have, want)
	}
}

func Test_answer_is_less_than_fifty(t *testing.T) {
	if have := answer(); have >= 50 {
		t.Errorf("answer: have %d; want: less than 50", have)
	}
}

func Test_answer_is_a_multiple_of_seven(t *testing.T) {
	if have, want := answer()%7, 0; have != want {
		t.Errorf("answer modulo 7: have %d; want: %d", have, want)
	}
}
