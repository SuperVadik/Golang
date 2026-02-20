package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/menu"
	"demo/password/output"
)

func main() {
	vaultWithDb := account.NewVault(files.NewJsonBd("data.json"))
	for {
		menuItem, err := menu.GetMenu()
		if err != nil {
			output.PrintError(err)
			continue
		}
		switch menuItem {
		case "create":
			createAccount(vaultWithDb)
		case "find":
			findAccount(vaultWithDb)
		case "delete":
			deleteAccount(vaultWithDb)
		case "exit":
			return
		default:
			output.PrintError("Неверный выбор, попробуйте снова.")
			continue
		}
	}
}
func createAccount(vaultWithDb *account.VaultWithDb) {

	login := output.PromtData([]string{"Введите логин: "})
	password := output.PromtData([]string{"Введите пароль: "})
	url := output.PromtData([]string{"Введите URL: "})

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Ошибка создания аккаунта: " + err.Error())
		return
	}

	vaultWithDb.AddAccount(*myAcc)
}

func findAccount(vaultWithDb *account.VaultWithDb) {
	searchStr := output.PromtData([]string{"Введите логин или URL для поиска: "})
	accList, err := vaultWithDb.FindAccount(searchStr)
	if err != nil {
		output.PrintError(err)
		return
	}
	for _, acc := range *accList {
		acc.Output()
	}
}

func deleteAccount(vaultWithDb *account.VaultWithDb) {
	searchStr := output.PromtData([]string{"Введите логин или URL для поиска: "})
	isDeleted, err := vaultWithDb.DeleteAccount(searchStr)
	if err != nil {
		output.PrintError("Ошибка при удалении аккаунта: " + err.Error())
		return
	}
	if isDeleted {
		output.PrintError("Аккаунт успешно удалён")
	}
}
