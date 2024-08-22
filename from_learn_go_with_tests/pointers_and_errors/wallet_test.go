package pointers_and_errors

import (
	"testing"
)

func assertBalance(wallet Wallet, t testing.TB, want Bitcoin) {
	t.Helper() // just for showing explicit error when test failed
	got := wallet.Balance()

	if got != want {
		// Note: %s to let 'fmt' call String() method of Bitcoin
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		// Act
		wallet.Deposit(10)

		want := Bitcoin(10)
		assertBalance(wallet, t, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		// Act
		wallet.Withdraw(Bitcoin(10))

		assertBalance(wallet, t, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(wallet, t, startingBalance)

		assertError(t, err, ErrInsufficientFunds.Error())
	})
}
