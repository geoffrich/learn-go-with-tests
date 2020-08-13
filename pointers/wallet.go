package wallet

import (
	"errors"
	"fmt"
)

// create new type from existing type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// lowercase symbol -- private outside of the package it's defined in
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// take a pointer to wallet so we can update it
	// otherwise a copy is passed
	// struct pointers are automatically dereferenced
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
