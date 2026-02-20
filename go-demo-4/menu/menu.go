package menu

import (
	"demo/password/output"
	"fmt"
)

func GetMenu() (string, error) {
	printMenu()
	key := readKey()
	switch key {
	case 1:
		return "create", nil
	case 2:
		return "find", nil
	case 3:
		return "delete", nil
	case 4:
		return "exit", nil
	default:
		return "", fmt.Errorf("Неверный ввод, попробуйте снова.")
	}
}

func printMenu() {
	promtData := []string{
		"____Приложение паролей____",
		"1: Создать аккаунт",
		"2: Найти аккаунт",
		"3: Удалить аккаунт",
		"4: Выход",
		"Выберите вариант",
	}
	output.PromtData(promtData)
}

func readKey() int {
	var key int
	fmt.Scan(&key)
	return key
}
