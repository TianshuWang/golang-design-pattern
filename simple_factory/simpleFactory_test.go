package simple_factory

import "testing"

func TestTypeHi(t *testing.T) {
	hiAPI, _ := NewAPI("HI")
	s := hiAPI.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("Type HI test fail")
	}
}

func TestTypeHello(t *testing.T) {
	helloAPI, _ := NewAPI("HELLO")
	s := helloAPI.Say("Tom")
	if s != "Hello, Tom" {
		t.Fatal("Type HELLO test fail")
	}
}

func TestTypeWrong(t *testing.T) {
	_, err := NewAPI("NADA")
	if err.Error() != "wrong api type passed" {
		t.Fatal("Wrong type test fail")
	}
}
