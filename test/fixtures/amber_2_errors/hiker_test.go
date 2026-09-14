package hiker

import "testing"

func Test_life_the_universe_and_everything(t *testing.T) {
	if have, want := answer(), 42; have != want {
		t.Errorf("answer: have %d; want: %d", have, want)
	}
}

func Test_digits_of_the_answer_sum_to_six(t *testing.T) {
	if have, want := answerDigitSum(), 6; have != want {
		t.Errorf("answerDigitSum: have %d; want: %d", have, want)
	}
}
