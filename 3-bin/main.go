package main

import (
	"bin/app-1/api"
	"bin/app-1/bins"
	"bin/app-1/file"
	"bin/app-1/storage"
)

func main() {
	bin := bins.NewBin("12345", true, "MyFirstBin")
	storage.SaveData()
	api.StartAPI()
	file.ReadFile()
	println("Bin created with ID:", bin.Id)
}
