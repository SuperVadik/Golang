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
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("invalid login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	acc := &Account{
		Login:     login,
		Url:       urlString,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		acc.generatePassword(12)
	}
	return acc, nil
}

func FindAccount(login, url string) (*Account, error) {
	return nil, nil
}

func findAccountByLogin(login string) (*Account, error) {
	return nil, nil
}

func findAccountByUrl(url string) (*Account, error) {
	return nil, nil
}

func DeleteAccount(login, url string) error {
	return nil
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login)
	//fmt.Println(acc.login, acc.password, acc.url)

}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}
	acc.Password = string(res)
}
