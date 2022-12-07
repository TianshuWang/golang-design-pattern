package state

import "testing"

func TestVendingMachine(t *testing.T) {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil && err.Error() != "item requested" {
		t.Fatal("Vending machine has item state fail")
	}

	err = vendingMachine.insertMoney(10)
	if err != nil && err.Error() != "money entered is ok" {
		t.Fatal("Vending machine item requested state fail")
	}

	err = vendingMachine.dispenseItem()
	if vendingMachine.currentState != vendingMachine.noItem {
		t.Fatal("Vending machine no item state fail")
	}

	err = vendingMachine.addItem(2)
	if vendingMachine.itemCount != 2 {
		t.Fatal("Vending machine add item fail")
	}

	err = vendingMachine.requestItem()
	err = vendingMachine.insertMoney(9)
	if err != nil && err.Error() != "inserted money is less. please insert $10" {
		t.Fatal("Vending machine insert money fail")
	}
}
