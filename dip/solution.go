package main

import (
	"errors"
	"time"
)

type PaymentMethod interface {
	Validate() error
	MakePayment()
}

type CreditCard struct {
	Number   string
	CVV      string
	SelfLife *time.Time
}

func (c *CreditCard) Validate() error {
	if c.Number == "" {
		return errors.New("number is required")
	}

	if c.CVV == "" {
		return errors.New("cvv is required")
	}

	if c.SelfLife == nil {
		return errors.New("validate is required")
	}

	return nil
}

func (c *CreditCard) MakePayment() {
	// pay
}

type MetaMask struct {
	Connected bool
}

func (cw *MetaMask) Validate() error {
	if !cw.Connected {
		return errors.New("you need to connect with your wallet")
	}

	return nil
}

func (cw *MetaMask) MakePayment() {
	// pay
}

// Pay depends on abstraction and not implementation
// thus making it extensible (OCP)
func Pay(paymentMethod PaymentMethod) error {
	if err := paymentMethod.Validate(); err != nil {
		return err
	}

	paymentMethod.MakePayment()

	return nil
}

func main() {
	creditCard := new(CreditCard)
	metaMask := &MetaMask{
		Connected: true,
	}

	Pay(creditCard)
	Pay(metaMask)
}
