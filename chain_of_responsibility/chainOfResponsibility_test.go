package chain_of_responsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "Tom"}

	reception.execute(patient)
}
