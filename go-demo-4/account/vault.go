package account

import (
	"demo/password/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewVault() *Vault {
	vault, err := getVault()
	if err != nil {
		color.Red(err.Error())
		return getNewVault()
	}

	return vault
}

func getNewVault() *Vault {
	return &Vault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Ошибка при преобразовании данных: %v", err)
		return
	}
	files.WriteFile(data, "data.json")
}

func (vault *Vault) DeleteAccount(searchStr string) (bool, error) {
	accList, err := vault.FindAccount(searchStr)
	if err != nil {
		return false, err
	}
	vault.Accounts = *difference(&vault.Accounts, accList)
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Ошибка при преобразовании данных: %v", err)
		return false, err
	}
	files.WriteFile(data, "data.json")
	return true, nil
}

func (vault *Vault) FindAccount(searchStr string) (*[]Account, error) {
	vault, err := getVault()
	if err != nil {
		return nil, fmt.Errorf("Ошибка при поиске аккаунта:", err)
	}

	var accList []Account
	for _, acc := range vault.Accounts {
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

func getVault() (*Vault, error) {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении файла: %v", err)
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке данных: %v", err)
	}

	return &vault, nil
}

func difference(slice1, slice2 *[]Account) *[]Account {
	toRemove := make(map[Account]bool)
	for _, item := range *slice2 {
		toRemove[item] = true
	}

	var result []Account
	for _, item := range *slice1 {
		if !toRemove[item] {
			result = append(result, item)
		}
	}
	return &result
}
