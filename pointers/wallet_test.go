package pointers

import "testing"

func TestWaller(t *testing.T) {

	t.Run("Depositing", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		want := Ethereum(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdrawing", func(t *testing.T) {
		wallet := Wallet{balance: Ethereum(15)}
		wallet.Withdraw(10)
		want := Ethereum(5)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdrawing too much", func(t *testing.T) {
		originalBalance := Ethereum(100)
		wallet := Wallet{originalBalance}
		err := wallet.Withdraw(999)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, originalBalance)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Ethereum) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("Expected error %v but got %v", want, got)
	}
}
