package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation := []string{"AWG", "SUM", "MED"}
	fmt.Println("Рассчёт массива")
	input, err := scan()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	operationVal := setOperation(operation)
	val, _ := calculate(input, operationVal)
	fmt.Printf("%s = %.2f", operationVal, val)
}

func calculate(numbers []float64, operationVal string) (float64, error) {
	sum := sum(numbers)

	switch operationVal {
	case "AWG":

		if len(numbers) == 0 {
			return 0, nil
		}
		if sum == 0 {
			return 0, nil
		}
		return sum / float64(len(numbers)), nil
	case "SUM":
		return sum, nil
	case "MED":
		sort.Float64s(numbers)
		length := len(numbers)
		if len(numbers)%2 == 0 {
			mid := length / 2
			return (numbers[mid-1] + numbers[mid]) / 2, nil
		}
		return numbers[length/2], nil

	default:
		return 0, fmt.Errorf("ошибка: неизвестная операция '%s'", operationVal)
	}
}

func sum(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func setOperation(operation []string) string {
	var val string
	for {
		fmt.Printf("Выберете операцию %s: ", strings.Join(operation, ", "))
		fmt.Scan(&val)
		val = strings.ToUpper(val)
		if slices.Contains(operation, val) {
			break
		}
		fmt.Print("Ошибка ввода! Введённое значение не соответствует списку")
	}
	return val
}

func scan() ([]float64, error) {
	var numArr []float64
	fmt.Print("Введите неограниченное число чисел через запятую (2, 10, 9): ")
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	strArr := strings.Split(strings.ReplaceAll(str, " ", ""), ",")
	for _, val := range strArr {
		num, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err != nil {
			return nil, fmt.Errorf("ошибка: '%s' не является числом", val)
		}
		numArr = append(numArr, num)
	}
	return numArr, err
}
