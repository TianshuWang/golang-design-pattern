package decoretor

type Component interface {
	Calculate() int
}

type ConcreteComponent struct {
}

func (*ConcreteComponent) Calculate() int {
	return 0
}

type MultiplyDecorator struct {
	Component
	num int
}

func (d *MultiplyDecorator) Calculate() int {
	return d.Component.Calculate() * d.num
}

func WrapMultiplyDecorator(c Component, num int) Component {
	return &MultiplyDecorator{
		Component: c,
		num:       num,
	}
}

type AddDecorator struct {
	Component
	num int
}

func (d *AddDecorator) Calculate() int {
	return d.Component.Calculate() + d.num
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}
