package leetcode2043

import "testing"

func TestBank(t *testing.T) {
	bank := Constructor([]int64{10, 100, 20, 50, 30})
	if !bank.Withdraw(3, 10) {
		t.Errorf("Withdraw failed")
	}

	if !bank.Transfer(5, 1, 20) {
		t.Errorf("Transfer failed")
	}

	if bank.Deposit(5, 20) != true {
		t.Errorf("Deposit failed")
	}

	if bank.Transfer(3, 4, 15) != false {
		t.Errorf("Transfer should have failed due to insufficient funds")
	}

	if bank.Withdraw(10, 50) != false {
		t.Errorf("Withdraw should have failed due to invalid account")
	}
}
