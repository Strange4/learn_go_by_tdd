package pointers

import (
	"errors"
	"fmt"
)

type Ethereum uint

func (e Ethereum) String() string {
	return fmt.Sprintf("%d ETH", e)
}

type Wallet struct {
	balance Ethereum
}

func (w *Wallet) Deposit(amount Ethereum) {
	w.balance += amount
}

func (w *Wallet) Balance() Ethereum {
	return w.balance
}

var ErrInsufficientFunds = errors.New("not enough funds to withdraw that amount")

func (w *Wallet) Withdraw(amount Ethereum) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
