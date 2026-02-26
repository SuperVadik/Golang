package storage

import (
	"bin/app-1/bins"
	"encoding/json"
)

type Storage interface {
	Read() ([]byte, error)
	Write(content []byte) error
}

func toByte(dataByte any) ([]byte, error) {
	file, err := json.Marshal(dataByte)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// SaveBins сохраняет список бинов
func SaveData(data any, db Storage) error {
	dataByte, err := toByte(data)
	if err != nil {
		return err
	}
	db.Write(dataByte)
	return nil
}

// LoadBins загружает список бинов
func LoadData(db Storage) (*[]bins.Bin, error) {
	dataByte, err := db.Read()
	if err != nil {
		return nil, err
	}
	var data []bins.Bin
	err = json.Unmarshal(dataByte, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
