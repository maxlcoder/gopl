package bank

import (
	"sync"
	"testing"
)

func TestWithdraw(t *testing.T) {
	Deposit(10000) // 存入 10000

	var wg sync.WaitGroup
	// 测试连续提现，总提现 0+1+2+...+99, n(n+1) , 99 * (99 +1) / 2 = 4950
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(amount int) {
			Withdraw(amount)
			wg.Done()
		}(i)
	}
	wg.Wait()
	// 检测10000 - 4950 = 5050
	if got, want := Balance(), 5050; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}