package main

import (
	"demo/password/account"
	"demo/password/files"
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

	myAcc, err := account.CreateAccount(login, password, url)

	file, err := myAcc.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать данные в JSON:", err)
		return
	}
	files.WriteFile(file, "data.json")
}

func promtData(promt string) string {
	fmt.Print(promt + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
