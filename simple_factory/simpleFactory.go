package simple_factory

import (
	"fmt"
	"strings"
)

type API interface {
	Say(name string) string
}

type hiAPI struct {
}

type helloAPI struct {
}

func NewAPI(t string) (API, error) {
	if strings.EqualFold(t, "HI") {
		return &hiAPI{}, nil
	} else if strings.EqualFold(t, "HELLO") {
		return &helloAPI{}, nil
	}

	return nil, fmt.Errorf("wrong api type passed")
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
