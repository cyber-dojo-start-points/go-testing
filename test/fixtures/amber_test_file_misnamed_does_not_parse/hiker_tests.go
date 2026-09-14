package hiker

import "testing"

func Test_answer_is_three_digits(t *testing.T) {
	if have := answer(); have < 100 {
