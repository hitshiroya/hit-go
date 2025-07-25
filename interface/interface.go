package main

import "fmt"

type paymenter interface {
	pay(amount int)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount int) {
	p.gateway.pay(amount)
}

type stripe struct{}

func (s stripe) pay(amount int) {
	fmt.Println("amount paid", amount)
}

func main() {

	stripeGw := stripe{}

	newPaymenter := payment{
		gateway: stripeGw,
	}

	newPaymenter.makePayment(1000)

}
