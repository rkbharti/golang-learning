package main

import (
	"fmt"
	"time"
)

// create interface as an contract because we use pay method for all the payment method
type PaymentMethod interface {
	Pay(amount float64) (string, error)
}

// create strcuct for different paymnet gatway UPI, CreditCard, Paypal
type Upi struct {
	upid string
}
type CreditCard struct {
	cardNumber string
	expiryDate string
	cvv        string
}
type Paypal struct {
	email string
}

// implemnet  method for each struct for pay -UPI
func (u *Upi) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", fmt.Errorf("invalid amount")
	}
	return fmt.Sprintf("Payment successful | TxnID: %d", time.Now().Unix()), nil
}

// for credit card
func (c *CreditCard) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", fmt.Errorf("Invalid Amount")
	}
	return fmt.Sprintf("Payment successful | TxnID: %d", time.Now().Unix()), nil
}

func (p *Paypal) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", fmt.Errorf("Invalid amount")
	}
	return fmt.Sprintf("Payment successful | TxnID: %d", time.Now().Unix()), nil
}

// use polymorphism to processPayment
func processPayment(p PaymentMethod, amount float64) {

	msg, err := p.Pay(amount)
	if err != nil {
		fmt.Println("Payment failed", err)
		return
	}
	fmt.Println(msg)

}

func main() {
	for {
		fmt.Println("learning interface by making a Payment System ")
		var choice int
		fmt.Println("---------Payment Method--------")
		fmt.Println("1: Upi")
		fmt.Println("2:Credit-Card")
		fmt.Println("3:Paypal")
		fmt.Println("4:Exit from Payemnt gateway")
		var amount float64
		fmt.Println("Enter Your Amount: ")
		fmt.Scanln(&amount)
		fmt.Println("Enter Your Choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Pay using UPI")
			var details = &Upi{}
			fmt.Println("Enter ur UPI ID")
			fmt.Scanln(&details.upid)
			processPayment(details, float64(amount))
		case 2:
			fmt.Println("Pay using credit card")
			var details = &CreditCard{}
			fmt.Println("Enter ur card number")
			fmt.Scanln(&details.cardNumber)
			fmt.Println("enter ur expiry date : mm/yy")
			fmt.Scanln(&details.expiryDate)
			fmt.Println("Enter ur cvv")
			fmt.Scanln(&details.cvv)
			fmt.Println("Card ending with:", details.cardNumber[len(details.cardNumber)-4:])
			processPayment(details, float64(amount))
		case 3:
			fmt.Println("entre ur email address to pay through PayPal")
			var email = &Paypal{}
			fmt.Println("Enter Ur email id for Paypal")
			fmt.Scanln(&email.email)
			processPayment(email, float64(amount))
		case 4:
			fmt.Println("Exiting from payment gateway")
			return
		default:
			fmt.Println("Invalid input")

		}

	}
}
