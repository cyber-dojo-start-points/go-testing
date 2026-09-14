package hiker

import "testing"

func Test_answer_is_two_digits(t *testing.T) {
	if have := answer(); have < 10 || have > 99 {
		t.Errorf("answer: have %d; want: two digits", have)
	}
}

func Test_answer_is_three_digits(t *testing.T) {
	if have := answer(); have < 100 || have > 999 {
		t.Errorf("answer: have %d; want: three digits", have)
	}
}
