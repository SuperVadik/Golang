package main

import (
	fmt "fmt"
	"slices"
	"strings"
)

/*
Функция входа в приложение
*/
func main() {
	currArr := []string{"RUB", "USD", "EUR"}
	currencyMap := initCurrency()
	fmt.Println("Калькулятор валют")
	orignCurrName := readCurrencyName("Введите исходную валюту: ", currArr)
	currArr = changeArray(currArr, orignCurrName)
	orignCurrNumber := readCurrencyNumber()
	targetCurrName := readCurrencyName("Введите целевую валюту: ", currArr)
	val := calculate(&currencyMap, orignCurrNumber, strings.ToUpper(orignCurrName), strings.ToUpper(targetCurrName))
	fmt.Printf("Результат: %.2f %s", val, targetCurrName)
}

/*
Инициализация курсов валют
*/
func initCurrency() map[string]map[string]float64 {
	var currencyMap = make(map[string]map[string]float64)
	currencyMap["USD"] = map[string]float64{"EUR": 0.87, "RUB": 81.25}
	currencyMap["EUR"] = map[string]float64{"EUR": 1.15, "RUB": 93.71}
	currencyMap["RUB"] = map[string]float64{"EUR": 0.011, "USD": 0.012}
	return currencyMap
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
		_, _ = fmt.Scan(&val)
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
		_, _ = fmt.Scan(&val)
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
func calculate(currencyMap *map[string]map[string]float64, number float64, origCurrStr, targetCurrStr string) float64 {

	origCurr := (*currencyMap)[origCurrStr]
	return origCurr[targetCurrStr] * number
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
