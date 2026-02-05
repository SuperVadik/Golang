package account

import (
	"demo/password/files"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Voult struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (voult *Voult) ToBytes() ([]byte, error) {
	file, err := json.Marshal(voult)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewVoult() *Voult {
	voult, err := getVoult()
	if err != nil {
		color.Red(err.Error())
		return getNewVault()
	}

	return voult
}

func getNewVault() *Voult {
	return &Voult{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}
}

func (voult *Voult) AddAccount(acc Account) {
	voult.Accounts = append(voult.Accounts, acc)
	voult.UpdatedAt = time.Now()
	data, err := voult.ToBytes()
	if err != nil {
		color.Red("Ошибка при преобразовании данных: %v", err)
		return
	}
	files.WriteFile(data, "data.json")
}

func (voult *Voult) RemoveAccount(index int) {
	if index < 0 || index >= len(voult.Accounts) {
		return
	}
	voult.Accounts = append(voult.Accounts[:index], voult.Accounts[index+1:]...)
	voult.UpdatedAt = time.Now()
}

func getVoult() (*Voult, error) {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении файла: %v", err)
	}

	var voult Voult
	err = json.Unmarshal(file, &voult)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке данных: %v", err)
	}

	return &voult, nil
}
