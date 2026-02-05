package menu

import "fmt"

func GetMenu() (string, error) {
	for {
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
			fmt.Println("Неверный ввод, попробуйте снова.")
			continue
		}
	}
}

func printMenu() {
	fmt.Println("____Приложение паролей____")
	fmt.Println("Выберите вариант:")
	fmt.Println("1: Создать аккаунт")
	fmt.Println("2: Найти аккаунт")
	fmt.Println("3: Удалить аккаунт")
	fmt.Println("4: Выход")
}

func readKey() int {
	var key int
	fmt.Scan(&key)
	return key
}
