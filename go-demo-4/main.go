package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*!@#$%^&()")

// account struct and methods
type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	account
	createdAt time.Time
	updetedAt time.Time
}

/*func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("invalid login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	acc := &account{
		login:    login,
		url:      urlString,
		password: password,
	}
	if password == "" {
		acc.generatePassword(12)
	}
	return acc, nil
}*/

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("invalid login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	acc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updetedAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" {
		acc.generatePassword(12)
	}
	return acc, nil
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)

}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}
	acc.password = string(res)
}

//

// main function
func main() {
	login := promtData("Enter your login:")
	password := promtData("Enter your password:")
	url := promtData("Enter the URL:")

	myAcc, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Не верный формат URL: ", err)
		return
	}
	myAcc.outputPassword()
}

// promtData function to get user input
func promtData(promt string) string {
	fmt.Print(promt + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
