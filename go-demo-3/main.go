package main

import (
	"fmt"
)

var bookmarks map[string]string

/*
Создать приложение, которое сначала выдаёт меню:
 - 1. Посмотреть закладки
 - 2. Добавить закладку
 - 3. Удалить закладку
 - 4. Выход

При 1 - Выводит закладки
При 2 - 2 поля ввода названия и адреса и после добавление
При 3 - Ввод названия и удаление по нему
При 4 - Завершение
*/

func main() {
	bookmarks = make(map[string]string)
menu:
	for {
		getMenu()
		val := scanData("Введите номер операции")
		switch val {
		case "1":
			findAllBookmarks()
			continue
		case "2":
			addBookmark()
			continue
		case "3":
			delBookmark()
			continue
		case "4":
			break menu
		}
	}
}

func getMenu() {
	fmt.Println(`1. Посмотреть закладки
2. Добавить закладку
3. Удалить закладку
4. Выход`)
}

func scanData(str string) string {
	fmt.Printf("%s: ", str)
	fmt.Scan(&str)
	return str
}

func addBookmark() {
	var val, key string
	key = scanData("Введите наименование")
	_, err := findBookmark(key)
	if err != nil {
		val = scanData("Введите адрес")
	} else {
		fmt.Println("Закладка с ключём %s уже существует", key)
		return
	}
	fmt.Println("Закладка добавлена")

	bookmarks[key] = val

}

func delBookmark() {
	key := scanData("Введите наименование")
	_, err := findBookmark(key)
	if err == nil {
		delete(bookmarks, key)
		fmt.Println("Закладка удалена")
	}
}

func findAllBookmarks() {
	if len(bookmarks) <= 0 {
		fmt.Println("Закладки отстутсвуют")
		return
	}
	for key := range bookmarks {
		bookmarkVal, err := findBookmark(key)
		if err == nil {
			printBookmark(key, bookmarkVal)
		}
	}
}

func findBookmark(key string) (string, error) {
	bookmarkVal := bookmarks[key]
	if bookmarkVal == "" {
		err := fmt.Errorf("Закладка с ключём %s не найдена", key)
		return "", err
	}
	return bookmarkVal, nil
}

func printBookmark(key, bookmarkVal string) {
	fmt.Println(key, bookmarkVal)
}
