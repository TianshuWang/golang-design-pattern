package decoretor

import "testing"

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WrapAddDecorator(c, 10)
	c = WrapMultiplyDecorator(c, 8)

	if res := c.Calculate(); res != 80 {
		t.Fatal("Test decorator fail")
	}
}
