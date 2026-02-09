package main

import (
	"bin/app-1/bins"
	"bin/app-1/storage"
	"fmt"
)

func main() {
	binList := bins.NewBinList()

	err := storage.SaveData(binList, "data.json")
	if err != nil {
		panic(err)
	}

	loadedData, err := storage.LoadData[[]bins.Bin](nil, "data.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Загруженные данные: %+v\n", *loadedData)
}
