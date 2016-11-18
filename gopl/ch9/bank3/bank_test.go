package bank_test

import (
	"sync"
	"testing"

	bank "github.com/ajaxhe/golang/gopl/ch9/bank3"
)

func TestBank(t *testing.T) {
	var n sync.WaitGroup

	for i := 0; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)

	}
	n.Wait()

	for i := 0; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Withdraw(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank.Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
