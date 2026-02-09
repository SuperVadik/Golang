package main

import (
	"demo/password/account"
	"demo/password/menu"
	"fmt"
)

func main() {
	vault := account.NewVault()
	for {
		menuItem, err := menu.GetMenu()
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch menuItem {
		case "create":
			createAccount(vault)
		case "find":
			findAccount(vault)
		case "delete":
			deleteAccount(vault)
		case "exit":
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
			continue
		}
	}
}
func createAccount(vault *account.Vault) {

	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите URL: ")

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Ошибка создания аккаунта:", err)
		return
	}

	vault.AddAccount(*myAcc)
}

func findAccount(vault *account.Vault) {
	searchStr := promtData("Введите логин или URL для поиска: ")
	accList, err := vault.FindAccount(searchStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, acc := range *accList {
		acc.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	searchStr := promtData("Введите логин или URL для поиска: ")
	isDeleted, err := vault.DeleteAccount(searchStr)
	if err != nil {
		fmt.Println("Ошибка при удалении аккаунта:", err)
		return
	}
	if isDeleted {
		fmt.Println("Аккаунт успешно удалён")
	}
}

func promtData(promt string) string {
	fmt.Print(promt + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
