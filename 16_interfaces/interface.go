package main

import "fmt"

// Defining an interface "paymenter"
// This interface declares a method "pay" which must be implemented by any struct that wants to be a payment gateway
type paymenter interface {
	pay(amount float32)
}

// The "payment" struct holds a reference to a payment gateway that implements the "paymenter" interface
type payment struct {
	gateway paymenter
}

// Struct for Razorpay payment gateway
type razorpay struct{}

// Implementing the "pay" method for Razorpay
type stripe struct{}

func (r razorpay) pay(amount float32) {
	// Logic to make a payment using Razorpay
	fmt.Println("making payment using razorpay", amount)
}

// Struct for Stripe payment gateway
func (s stripe) pay(amount float32) {
	// Logic to make a payment using Stripe
	fmt.Println("making payment using stripe", amount)
}

// Struct for PayPal payment gateway
type paypal struct{}

// Implementing the "pay" method for PayPal
func (p paypal) pay(amount float32) {
	// Logic to make a payment using PayPal
	fmt.Println("making payment using paypal", amount)
}

// Method to make a payment using the specified payment gateway
func (p payment) makePayment(amount float32) {
	// Calling the "pay" method on the assigned payment gateway
	p.gateway.pay(amount)
}

func main() {
	// stripePaymentGw := stripe{}
	// razorPaymentGw:= razorpay{}

	// Creating an instance of PayPal as a payment gateway
	paypalGw := paypal{}

	// Creating a "payment" instance and assigning PayPal as the payment gateway
	newPayment := payment{
		gateway: paypalGw,
	}

	// Making a payment using PayPal
	newPayment.makePayment(100)
}

/*
Why Use Interfaces in Go?
1. **Decoupling**: Interfaces allow us to write flexible and reusable code by abstracting implementations.
2. **Dependency Injection**: We can inject any payment gateway (e.g., Razorpay, Stripe, PayPal) without modifying existing logic.
3. **Polymorphism**: Different payment methods can be handled uniformly using the same function calls.
4. **Testing**: Interfaces help in writing mock implementations for unit tests.

How to Use Interfaces in Go?
- Define an interface with the required methods.
- Implement the interface in different structs.
- Use an interface type in a struct or function to accept any implementation.
- Call the methods on the interface to achieve runtime polymorphism.
*/
