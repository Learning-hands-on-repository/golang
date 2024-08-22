package pointers_and_errors

import (
	"errors"
	"fmt"
)

// Good for adding domain specific to existing types
type Bitcoin int

type Stringer interface {
	String()  string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// this won't works since 'wallet' is copied, not same address
// func (wallet Wallet) Deposit (amount int) {
// 	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
// 	wallet.balance += amount
// }

// this won't works since 'wallet' is copied, not same address
// func (wallet Wallet) Balance () int {
// 	return 0
// }

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet Wallet) Balance() Bitcoin {
	return wallet.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) (error) {
	if (wallet.balance < amount) {
		return ErrInsufficientFunds
	}
	
	wallet.balance -= amount
	return nil
}
