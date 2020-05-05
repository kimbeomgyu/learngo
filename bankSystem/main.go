package main

import (
	"fmt"
	"learngo/BankSystem/accounts"
	"log"
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
