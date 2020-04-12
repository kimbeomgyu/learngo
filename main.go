package main

import (
	"fmt"
	"log"

	"github.com/myProject/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kimbeomgyu")
	account.Deposit(10000)
	fmt.Println(account.Balance())
	err := account.Withdraw(10000)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account)
}
