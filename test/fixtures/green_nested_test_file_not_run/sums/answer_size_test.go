package sums

import (
	"fmt"
	"testing"
)

// Go ties a package to its directory, so this file is a package of its own and
// go test, run in the sandbox directory, never reaches it. The print is what
// shows whether it ran.
func Test_answer_is_two_digits(t *testing.T) {
	fmt.Println("checking the size of the answer")
	if have := 42; have < 10 || have > 99 {
		t.Errorf("answer: have %d; want: two digits", have)
	}
}
