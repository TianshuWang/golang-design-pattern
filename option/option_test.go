package option

import "testing"

func TestOptionDefault(t *testing.T) {
	server := NewServer()

	if server.Addr != "127.0.0.1" || server.Port != 8000 {
		t.Fatal("Test Option default fail")
	}
}

func TestOptionWithAttrPort(t *testing.T) {
	server := NewServer(WithAddr("localhost"), WithPort(8001))

	if server.Addr != "localhost" || server.Port != 8001 {
		t.Fatal("Test Option with addr and port fail")
	}
}
