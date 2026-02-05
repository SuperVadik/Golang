package main

import (
	"demo/password/account"
	"demo/password/menu"
	"fmt"
)

func main() {
	_, err := menu.GetMenu()
	if err != nil {
		fmt.Println("Ошибка меню:", err)
		return
	}
	createAccount()
}
func createAccount() {

	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите URL: ")

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Ошибка создания аккаунта:", err)
		return
	}
	vault := account.NewVoult()
	vault.AddAccount(*myAcc)
}

func promtData(promt string) string {
	fmt.Print(promt + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
