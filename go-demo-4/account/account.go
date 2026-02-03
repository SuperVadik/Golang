package account

import (
	"errors"
	//"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*!@#$%^&()")

// Account struct and methods
type Account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	Account
	createdAt time.Time
	updetedAt time.Time
}

func newAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("invalid login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	acc := &Account{
		login:    login,
		url:      urlString,
		password: password,
	}
	if password == "" {
		acc.generatePassword(12)
	}
	return acc, nil
}

func NewAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
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
		Account: Account{
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

func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
	//fmt.Println(acc.login, acc.password, acc.url)

}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}
	acc.password = string(res)
}
