package main

import (
	"fmt"
	"math"
	"strings"
)

const power float64 = 2

func main() {
	fmt.Println("Калькулятор индекса массы тела")
	for {
		userHeight, userKg := getUserInput()
		imt := calculateImt(userHeight, userKg)
		outputResult(imt)
		if !checkRepeatCalculation() {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш ИМТ: %.0f\n", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateImt(userHeight, userKg float64) float64 {
	return userKg / math.Pow(userHeight/100, power)
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите выполнить ещё расчёт? (y/n)")
	fmt.Scan(&userChoise)
	return strings.ToUpper(userChoise) == "Y"
}
