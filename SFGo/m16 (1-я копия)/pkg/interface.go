package pkg

import (
	"fmt"
)

type Client struct {
	deposit    int
	withdrawal int
	balance    int
}

type BankClient interface {
	Deposit(Amount int) (string, int)

	Withdrawal(Amount int) (string, int)

	Balance(Amount int) (string, int)
}

func (c *Client) Deposit(Amount int) (string, int) {
	fmt.Println("какую сумму хотите внести?")
	_, err := fmt.Scan(&c.deposit)
	checkerr(err)
	if c.deposit == 0 {
		return "Вы ничего не внесли", 0
	}
	if c.deposit >= 1 {
		return "Вы внесли деньги, теперь ваш баланс - ", Amount + c.deposit
	} else {
		return "Ошибка", 0
	}

}

func (c *Client) Withdrawal(Amount int) (string, int) {
	fmt.Println("Какую сумму хотите снять?")
	_, err := fmt.Scan(&c.withdrawal)
	checkerr(err)
	if c.withdrawal == 0 {
		return "Вы ничего не сняли", 0
	}
	if Amount >= c.withdrawal {
		return "С нятие наличных, доступный баланс - ", Amount - c.withdrawal
	}
	if Amount < c.withdrawal {
		return "Недостаточно средств.", 0
	} else {
		return "Ошибка", 0
	}
}

func (c *Client) Balance(Amount int) (string, int) {
	c.balance = Amount
	return "Ваш текущий баланс - ", c.balance
}

func Strcreate() *Client {
	return &Client{}
}
