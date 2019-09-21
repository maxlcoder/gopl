package bank

import "sync"

var (
	mu sync.Mutex // 保护 balance
	balance int
)

func Deposit(amount int) {
	mu.Lock() // 上锁
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

/*
// 注意：不是原子操作，并发执行的时候有存在竞争
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false // 余额不足
	}
	return true
}
 */

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if Balance() < 0 {
		deposit(amount)
		return false // 余额不足
	}
	return true
}

func deposit(amount int) {
	balance += amount
}