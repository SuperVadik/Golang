package storage

import (
	"bin/app-1/files"
	"encoding/json"
)

func toByte(dataByte any) ([]byte, error) {
	file, err := json.Marshal(dataByte)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// SaveBins сохраняет список бинов в файл bins.txt
func SaveData(data any, name string) error {
	dataByte, err := toByte(data)
	if err != nil {
		return err
	}
	files.WriteFile(dataByte, name)
	return nil
}

// LoadBins загружает список бинов из файла bins.txt
func LoadData[T any](data *T, name string) (*T, error) {
	dataByte, err := files.ReadFile(name)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataByte, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
