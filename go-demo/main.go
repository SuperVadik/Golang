package main

import (
	"fmt"
	"math"
)

const power float64 = 2

func main() {
	fmt.Println("Калькулятор индекса массы тела")
	userHeight, userKg := GetUserInput()
	imt := CalculateImt(userHeight, userKg)
	OutputResult(imt)
}

func OutputResult(imt float64) {
	fmt.Printf("Ваш ИМТ: %.0f", imt)
}

func CalculateImt(userHeight, userKg float64) float64 {
	return userKg / math.Pow(userHeight/100, power)
}
func GetUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}
