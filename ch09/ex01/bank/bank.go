package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan *Withdraws)

type Withdraws struct {
	amount int
	ch     chan bool
}

func Deposit(amount int) { deposits <- amount }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- &Withdraws{amount, ch}
	return <-ch
}

func Balance() int { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case wd := <-withdraws:
			if balance >= wd.amount {
				balance -= wd.amount
				wd.ch <- true
			} else {
				wd.ch <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
