package hiker

import (
	"fmt"
	"testing"
)

// go test prints nothing for a test that passes, so this print is what shows
// the file was collected and run.
func Test_answer_is_two_digits(t *testing.T) {
	fmt.Println("checking the size of the answer")
	if have := answer(); have < 10 || have > 99 {
		t.Errorf("answer: have %d; want: two digits", have)
	}
}
