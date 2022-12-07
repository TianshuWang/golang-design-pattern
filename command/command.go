package command

import "fmt"

type Command interface {
	Execute()
}

type MotherBoard struct {
}

type StartCommand struct {
	mb *MotherBoard
}

type RebootCommand struct {
	mb *MotherBoard
}

func (*MotherBoard) Start() {
	fmt.Print("System starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("System rebooting\n")
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (s *StartCommand) Execute() {
	s.mb.Start()
}

func (r *RebootCommand) Execute() {
	r.mb.Reboot()
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
