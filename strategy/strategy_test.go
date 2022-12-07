package strategy

import "testing"

func TestPayByCash(t *testing.T) {
	payment := NewPayment("Tom", "0001", 100, &Cash{})
	if payment.Pay() != "Pay $100 to Tom by cash" {
		t.Fatal("Payment strategy Cash test fail")
	}
}

func TestPayByBank(t *testing.T) {
	payment := NewPayment("Tom", "0002", 200, &Bank{})
	if payment.Pay() != "Pay $200 to Tom by bank account" {
		t.Fatal("Payment strategy Bank test fail")
	}
}
