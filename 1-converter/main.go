package main

import (
	"fmt"
	"slices"
	"strings"
)

type currency struct {
	orignCurr  string
	targetCurr string
	value      float64
}

/*
Функция входа в приложение
*/
func main() {
	currArr := []string{"RUB", "USD", "EUR"}
	fmt.Println("Калькулятор валют")
	orignCurrName := readCurrencyName("Введите исходную валюту: ", currArr)
	currArr = changeArray(currArr, orignCurrName)
	orignCurrNumber := readCurrencyNumber()
	targetCurrName := readCurrencyName("Введите целевую валюту: ", currArr)
	val := calculate(orignCurrNumber, strings.ToUpper(orignCurrName), strings.ToUpper(targetCurrName))
	fmt.Printf("Результат: %.2f %s", val, targetCurrName)
}

/*
Инициализация курсов валют
*/
func initCurrency() (currency, currency) {
	usdRub := currency{
		orignCurr:  "USD",
		targetCurr: "RUB",
		value:      83.37,
	}

	usdEur := currency{
		orignCurr:  "USD",
		targetCurr: "EUR",
		value:      0.85,
	}

	return usdRub, usdEur
}

/*
Чтение наименования валюты из консоли
*/
func readCurrencyName(text string, currArr []string) string {
	var val string
	for {

		currChoice := fmt.Sprintf("Валюта доступная для выбора: %s", strings.Join(currArr, ", "))
		fmt.Println(currChoice)
		fmt.Print(text)
		fmt.Scan(&val)
		val = strings.ToUpper(val)
		if slices.Contains(currArr, val) {
			break
		}
		fmt.Println("Ошибка ввода! Введённое значение не соответствует списку доступной валюты")
	}

	return val
}

/*
Чтение количества валюты из консоли
*/
func readCurrencyNumber() float64 {
	var val float64
	for {
		fmt.Print("Введите количество валюты: ")
		fmt.Scan(&val)
		if val > 0 {
			break
		}
		fmt.Println("Ошибка ввода! Введённое значение меньше нуля или не является числом")
	}

	return val
}

/*
Вычисление курса валют относительно друг друга
*/
func calculate(number float64, orignCurr, targetCurr string) float64 {
	var val float64
	usdRub, usdEur := initCurrency()

	switch orignCurr {
	case "USD":
		if targetCurr == "RUB" {
			val = usdRub.value * number
		} else {
			val = usdEur.value * number
		}
	case "EUR":
		if targetCurr == "RUB" {
			val = usdRub.value * (number / usdEur.value)
		} else {
			val = number / usdEur.value
		}
	case "RUB":
		if targetCurr == "USD" {
			val = number / usdRub.value
		} else {
			val = usdEur.value * (number / usdRub.value)
		}

	}

	return val
}

/*
Изменение списка
*/
func changeArray(currArr []string, target string) []string {
	index := -1
	for i, v := range currArr {
		if v == target {
			index = i
			break
		}
	}

	// Если элемент найден, удаляем его
	if index != -1 {
		currArr = append(currArr[:index], currArr[index+1:]...)
	}
	return currArr
}
