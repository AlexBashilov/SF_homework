package pkg

import (
	"fmt"
	"os"
)

var act string

func Action() (string, int) {
	newclient := Strcreate()
	_, err := fmt.Scan(&act)
	checkerr(err)

	switch {
	case act == "deposit":
		return newclient.Deposit(Amount)
	case act == "withdrawal":
		return newclient.Withdrawal(Amount)
	case act == "balance":
		return newclient.Balance(Amount)
	case act == "exit":
		return "Выход", 0
	default:
		return ("Unsupported command. You can use commands: balance, deposit, withdrawal, exit"), 0
	}
}

func checkerr(err error) {
	if err != nil {
		fmt.Println("FATAL: ", err.Error())
		os.Exit(1)
	}
}
