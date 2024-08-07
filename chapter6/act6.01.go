package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidLastName = errors.New("invalid last name")
	ErrInvalidRoutingNumner = errors.New("invalid routing number")
)

type directDeposit struct {
	lastName string
	firstName string
	bankName string
	routingNumber int
	accountNumber int
}

func (dd *directDeposit) report() {
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println("First Name: ", dd.firstName)
	fmt.Println("Last Name: ", dd.lastName)
	fmt.Println("Bank Name: ", dd.bankName)
	fmt.Println("Routing Number: ", dd.routingNumber)
	fmt.Println("Account Number: ", dd.accountNumber)

}

func (dd *directDeposit) validateRoutingNumber() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if dd.routingNumber < 100 {
		panic(ErrInvalidRoutingNumner)
	}

	return nil
}

func (dd *directDeposit) validateLastName() error {
	dd.lastName = strings.TrimSpace(dd.lastName)
	if len(dd.lastName) == 0 {
		return ErrInvalidLastName
	}

	return nil
}


func main() {
	dd := directDeposit{

		lastName:      "  ",
		firstName:     "Abe",
		bankName:      "WilkesBooth Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	err := dd.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}
	err = dd.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	dd.report()

}