package main

import "goworkspase/bank"

func main() {
	var account = &bank.Account{
		AccountNum: "12345",
		Password:   "123",
		Balance:    100,
	}
	account.Query("123")
	account.Deposited(100, "123")
	account.WithDraw(50, "123")
	account.Query("123")
}
