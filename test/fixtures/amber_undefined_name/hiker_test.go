package hiker

import "testing"

func Test_life_the_universe_and_everything(t *testing.T) {
	if have, want := answerToEverything(), 42; have != want {
		t.Errorf("answer: have %d; want: %d", have, want)
	}
}
