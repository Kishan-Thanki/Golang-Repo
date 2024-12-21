package main

import "fmt"

type Payment struct {
	gateway Stripe
}

func (P Payment) makePayment(amount float32) {
	P.gateway.pay(amount)
}

type RazorPay struct{}

func (R RazorPay) pay(amount float32) {
	fmt.Println("Making payment using razorpay", amount)
}

type Stripe struct{}

func (S Stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe", amount)
}

func main() {
	newPayment := Payment{}
	newPayment.makePayment(100)
}
