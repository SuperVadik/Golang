package main

import "fmt"

func main() {
	m := map[string]string{
		"PurpleScool": "purplescool.ru",
	}
	fmt.Println(m)
	fmt.Println(m["PurpleScool"])
	m["PurpleScool"] = "https://purpleschool.ru"
	fmt.Println(m)
	m["Google"] = "https://google.com"
	m["Yandex"] = "https://yandex.ru"
	fmt.Println(m)
	delete(m, "Yandex")
	fmt.Println(m)
}
