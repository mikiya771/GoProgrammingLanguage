package bank_test

import (
	"testing"

	"github.com/mikiya771/gpl/ch09/ex01/bank"
)

func TestWithdrawNormal(t *testing.T) {
	bank.Deposit(200)
	ok := bank.Withdraw(100)
	if !ok {
		t.Error("Result is false, want true")
		return
	}
	ok = bank.Withdraw(100)
	if !ok {
		t.Error("Result is false, want true")
		return
	}
	ok = bank.Withdraw(100)
	if ok {
		t.Error("Result is true, want false")
		return
	}
	if bank.Balance() != 0 {
		t.Errorf("Result is %d, want 0", bank.Balance())
	}
}
