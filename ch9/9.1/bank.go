// bank 包提供了一个只有一个账户的并发安全银行
package bank

var deposits = make(chan int)  // 发送存款额
var balances = make(chan int)  // 接收余额

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

type withdraw struct {
	amount int
	succeed chan bool
}

var withdraws = make(chan withdraw)

func Withdraw(amount int) bool {
	succeed := make(chan bool)
	withdraws <- withdraw{amount, succeed}
	return <-succeed
}

func teller() {
	var balance int // balance 被限制在 teller goroutine 中
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraws:
			if withdraw.amount <= balance {
				balance -= withdraw.amount
				withdraw.succeed <- true
			} else {
				withdraw.succeed <- false
			}
		}
	}
}

func init() {
	go teller() // 启动监控 goroutine
}
