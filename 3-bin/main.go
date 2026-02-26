package main

import (
	"bin/app-1/bins"
	"bin/app-1/files"
	"bin/app-1/storage"
	"fmt"
)

func main() {
	db := files.NewJsonDb("data.json")

	binList := bins.NewBinList()
	err := storage.SaveData(binList, db)
	if err != nil {
		panic(err)
	}

	loadedData, err := storage.LoadData(db)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Загруженные данные: %+v\n", *loadedData)
}
