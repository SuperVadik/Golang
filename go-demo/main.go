package main

import (
	"fmt"
	"math"
)

const power float64 = 2

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	fmt.Println("Калькулятор индекса массы тела")
	userHeight, userKg := getUserInput()
	imt := calculateImt(userHeight, userKg)

	outputResult(imt)
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

func outputResult(imt float64) {
	fmt.Printf("Ваш ИМТ: %.0f\n", imt)
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
