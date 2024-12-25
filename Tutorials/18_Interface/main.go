package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type Payment struct {
	gateway paymenter
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
	// stripePaymentGW := Stripe{}
	razorpayPaymentGW := RazorPay{}
	newPayment := Payment{
		// gateway: stripePaymentGW,
		gateway: razorpayPaymentGW,
	}
	newPayment.makePayment(100)
}
