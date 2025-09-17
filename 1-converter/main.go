package main

import "fmt"

func main() {
	const usdEur float64 = 0.85
	const usdRub float64 = 82.82

	eurRub := usdRub / usdEur
	fmt.Print(eurRub)
}

func readCurrency() string {
	var val string
	fmt.Scan(&val)
	return val
}

func calculate(number float64, orignCurr, targetCurr string) float64 {
	return 0
}
