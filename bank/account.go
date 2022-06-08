package bank

import "fmt"

type Account struct {
	AccountNum string
	Password   string
	Balance    float64
}

// Deposited 存款
func (account *Account) Deposited(money float64, password string) {
	// 1.判断密码是否正确
	if password != account.Password {
		fmt.Println("密码不正确！")
		return
	}
	// 2.判断金额是否正确
	if money <= 0 {
		fmt.Println("输入的金额不正确！")
		return
	}
	// 存款
	account.Balance += money
	fmt.Println("业务办理成功！")
}

// WithDraw 存款
func (account *Account) WithDraw(money float64, password string) {
	// 1.判断密码是否正确
	if password != account.Password {
		fmt.Println("密码不正确！")
		return
	}
	// 2.判断金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("输入的金额不正确！")
		return
	}
	// 存款
	account.Balance -= money
	fmt.Println("业务办理成功！")
}

// Query 查询余额
func (account *Account) Query(password string) {
	// 1.判断密码是否正确
	if password != account.Password {
		fmt.Println("密码不正确！")
		return
	}
	fmt.Printf("账户余额：%v \n", account.Balance)
}
