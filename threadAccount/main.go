package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	프로그램
	논리적인 APP
	프로세스
	cpu에서 실행되고 있는상태 / 여러개의 쓰레드를 가질 수 있다
	쓰레드
	각 프로세스 내부에서 수행되고 있는 하나의 작업
	메모리가 엉망진창이 되는 예시
	CPU에서 동시에 여러개의 동작이 실행되는 경우 sync 문제가 발생할 수 있다.
	예를 들면 하나의 도화지에 여러명의 어린이가 그림을 마구잡이로 그리는 형태 -> 하나의 메모리에 여러 쓰레드가 명령을 내리는 형태
	대표적인 해결방법 : Lock
	하나의 도화지에 여러명의 어린이가 그림을 그리는 대시 한명의 어린이가 그림을 그릴때에는 다른 어린이가 그림을 그리지 못하도록 접근을 막는 방법
	-> 메모리에 lock를 걸어 접근을 할 수 없게 하는 방법 / Mutex
	어떤 시점에 lock을 해줘야 하는지는 굉장히 복잡하다.
	또다른 방법 : channel
	go에서만 지원되는 방법 -> 완벽한 해결책은 아니다.


	------------------------------------------
	Lock을 사용하다가 발생할 수 잇는 문제: DeadLock
	lock하는 메모리를 서로 함께 잡는 순간 같이 멈춰버리는 증상
	실무에서 발생되면 고치기 어려운 문제가 발생한다
	DeadLock을 해결하는 방법
	1. lock의 범위를 작게 잡는 방법
		a, b, c를 각각 lock하고 unlock하고를 반복하는 방법
	2. lock의 범위를 크게 잡는 방법
		a, b, c를 한번에 lock하고 다음 a, b, c들은 기다리고 unlock이 되면 그때 다시 lock을 하는 방법
*/

// Account is Account
type Account struct {
	balance int // 메모리 영역
	mutex   *sync.Mutex
}

// Withdraw is Withdraw
func (a *Account) Withdraw(val int) {
	a.balance -= val
}

// Deposit is Deposit
func (a *Account) Deposit(val int) {
	a.balance += val
}

// Balance is Balance
func (a *Account) Balance() int {
	balance := a.balance
	return balance
}

var accounts []*Account
var globalLock *sync.Mutex

func transfer(sender, receiver int, money int) {
	globalLock.Lock()
	accounts[sender].Withdraw(money)
	accounts[receiver].Deposit(money)
	globalLock.Unlock()
	fmt.Println("Transfer,", sender, receiver, money)
}

func getTotalBalance() int {
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	return total
}

func randomTransfer() {
	var sender, balance int
	for {
		sender := rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	transfer(sender, receiver, money)

}

func goTransfer() {
	for {
		randomTransfer()
	}
}

func printTotalBalance() {
	fmt.Printf("Total: %d\n", getTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
	}

	globalLock = &sync.Mutex{}
	go func() {
		for {
			transfer(0, 1, 100)
		}
	}()

	go func() {
		for {
			transfer(1, 0, 100)
		}
	}()

	for i := 0; i < 10; i++ {
		go goTransfer()
	}

	for {

		time.Sleep(100 * time.Millisecond)
	}
}
