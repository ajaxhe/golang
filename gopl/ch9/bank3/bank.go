package bank

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func deposit(amount int) {
	balance += amount
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
