package strategy

import "fmt"

type PaymentContext struct {
	Name   string
	CardID string
	Money  int
}

type PaymentStrategy interface {
	Pay(context *PaymentContext) string
}

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type Cash struct {
}

type Bank struct {
}

func NewPayment(name, cardID string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardID,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() string {
	return p.strategy.Pay(p.context)
}

func (*Cash) Pay(ctx *PaymentContext) string {
	return fmt.Sprintf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

func (*Bank) Pay(ctx *PaymentContext) string {
	return fmt.Sprintf("Pay $%d to %s by bank account", ctx.Money, ctx.Name)
}
