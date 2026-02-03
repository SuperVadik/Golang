package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	login := promtData("Enter your login:")
	password := promtData("Enter your password:")
	url := promtData("Enter the URL:")

	myAcc, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Не верный формат URL: ", err)
		return
	}
	myAcc.OutputPassword()
	files.WriteFile()
	files.ReadFile()
}

func promtData(promt string) string {
	fmt.Print(promt + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
