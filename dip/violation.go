//go:build violation
// +build violation

package main

import (
	"errors"
	"time"
)

type CreditCard struct {
	Number   string
	CVV      string
	SelfLife *time.Time
}

func Pay(creditCard *CreditCard) error {
	// this is not legal because payment depends on a credit card implementation
	// and if in the future we need the Pay function to execute another type of payment?
	// we will end up breaking the OCP, SRP

	if creditCard.Number == "" {
		return errors.New("number is required")
	}

	if creditCard.CVV == "" {
		return errors.New("cvv is required")
	}

	if creditCard.SelfLife == nil {
		return errors.New("validate is required")
	}

	// make payment

	return nil
}

func main() {
	creditCard := new(CreditCard)

	Pay(creditCard)
}
