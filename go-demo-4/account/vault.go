package account

import (
	"demo/password/output"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write(content []byte) error
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewVault(db Db) *VaultWithDb {
	vault, err := getVault(db)
	if err != nil {
		output.PrintError(err.Error())
		return getNewVault(db)
	}
	return vault
}

func getNewVault(db Db) *VaultWithDb {
	return &VaultWithDb{
		Vault: Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		},
		db: db,
	}
}

func (vaultWithDb *VaultWithDb) AddAccount(acc Account) {
	vaultWithDb.Accounts = append(vaultWithDb.Accounts, acc)
	vaultWithDb.save()
}

func (vaultWithDb *VaultWithDb) DeleteAccount(searchStr string) (bool, error) {
	accList, err := vaultWithDb.FindAccount(searchStr)
	if err != nil {
		return false, err
	}
	vaultWithDb.Accounts = *difference(&vaultWithDb.Accounts, accList)
	vaultWithDb.save()
	return true, nil
}

func (vaultWithDb *VaultWithDb) FindAccount(searchStr string) (*[]Account, error) {

	var accList []Account
	for _, acc := range vaultWithDb.Vault.Accounts {
		if strings.Contains(acc.Login, searchStr) || strings.Contains(acc.Url, searchStr) {
			accList = append(accList, acc)
		}
	}
	return &accList, nil
}

func getVault(db Db) (*VaultWithDb, error) {
	file, err := db.Read()
	if err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке данных: %v", err)
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке данных: %v", err)
	}
	vaultWithDb := VaultWithDb{
		Vault: vault,
		db:    db,
	}
	return &vaultWithDb, nil
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

func (vaultWithDb *VaultWithDb) save() {
	vaultWithDb.UpdatedAt = time.Now()
	data, err := vaultWithDb.ToBytes()
	if err != nil {
		output.PrintError("Ошибка при преобразовании данных: " + err.Error())
		return
	}
	vaultWithDb.db.Write(data)
}
