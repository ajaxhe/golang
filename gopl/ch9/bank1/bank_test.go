package bank_test

import (
	"fmt"
	"testing"

	//bank "gopl.io/ch9/bank1"
	bank "github.com/ajaxhe/golang/gopl/ch9/bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	go func() {
		bank.Deposit(200)
		fmt.Printf("=%d\n", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
