package main

import "fmt"

// В цикле спрашивать ввод транзакций: -10, 10, 40.5
// Добавлять каждую в массив транзакций
// Вывести массив
// Вывести сумму баланса в консоль

func main() {

	res := make([]int, 6, 6)

	fmt.Println(len(res), cap(res))

	/*
	   transactions := []float64{}

	   	for {
	   		transaction := scanTransaction()
	   		if transaction == 0 {
	   			break
	   		}
	   		transactions = append(transactions, transaction)
	   	}

	   fmt.Printf("Ваш баланс: %.2f", calculateBalans(transactions))
	*/
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalans(transactions []float64) float64 {
	balans := 0.0
	for _, val := range transactions {
		balans += val
	}
	return balans
}
