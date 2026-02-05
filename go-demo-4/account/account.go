package account

import (
	"errors"
	"strings"

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

func FindAccount(searchStr string) (*[]Account, error) {
	voult, err := getVoult()
	if err != nil {
		return nil, err
	}

	var accList []Account
	for _, acc := range voult.Accounts {
		if strings.Contains(acc.Login, searchStr) || strings.Contains(acc.Url, searchStr) {
			accList = append(accList, acc)
		}
	}
	return &accList, nil
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

func (acc *Account) Output() {
	color.Cyan(acc.Login)
	color.Green(acc.Password)
	color.Blue(acc.Url)
	color.Yellow("Дата создания: %s", acc.CreatedAt.Format("2006-01-02 15:04:05"))
	color.Magenta("Дата обновления: %s", acc.UpdatedAt.Format("2006-01-02 15:04:05"))
	color.White("--------------------")

}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}
	acc.Password = string(res)
}
