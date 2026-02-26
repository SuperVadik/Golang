package files

import (
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	name string
}

func NewJsonDb(name string) *File {
	return &File{name: name}
}

func (f File) Read() ([]byte, error) {
	ext := "json"
	isTrueExt := checkFileExtension(f.name, ext)
	if !isTrueExt {
		return nil, fmt.Errorf("%s", "неподдерживаемый формат файла. Ожидается "+"."+ext)
	}
	data, err := os.ReadFile(f.name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (f File) Write(content []byte) error {
	file, err := os.Create(f.name)
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

// проверяет расширение файла на .json
func checkFileExtension(name, ext string) bool {
	filepath.Ext(name)
	if filepath.Ext(name) != "."+ext {
		return false
	}
	return true
}
