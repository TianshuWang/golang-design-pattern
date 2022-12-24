package decoretor

import "testing"

func TestHigherOrder(t *testing.T) {
	a := 2
	b := 8
	decorator := execTime(Multiply)
	if c := decorator(a, b); c != 16 {
		t.Fatal("Test higher order fail")
	}
}
