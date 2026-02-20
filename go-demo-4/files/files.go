package files

import (
	"fmt"
	"os"
)

type JsonBd struct {
	fileName string
}

func NewJsonBd(name string) *JsonBd {
	return &JsonBd{
		fileName: name,
	}
}

func (db *JsonBd) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonBd) Write(content []byte) error {
	file, err := os.Create(db.fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Запись успешна")
	return nil
}
