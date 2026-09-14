package hiker

import "testing"

func Test_answer_is_two_digits(t *testing.T) {
	if have := answer(); have < 10 {
