package main

import (
	"demo/password/account"
	"demo/password/menu"
	"fmt"
)

func main() {
	for {
		menuItem, err := menu.GetMenu()
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch menuItem {
		case "create":
			createAccount()
		case "find":
			findAccount()
		case "delete":
			fmt.Println("Удалить аккаунт")
		case "exit":
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
			continue
		}
	}
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

func findAccount() {
	searchStr := promtData("Введите логин или URL для поиска: ")
	accList, err := account.FindAccount(searchStr)
	if err != nil {
		fmt.Println("Ошибка при поиске аккаунта:", err)
		return
	}
	if accList == nil || len(*accList) == 0 {
		fmt.Println("Аккаунт не найден")
		return
	}
	for _, acc := range *accList {
		acc.Output()
	}
}

func promtData(promt string) string {
	fmt.Print(promt + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
