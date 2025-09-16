package main

import "fmt"

func main() {
	const usdEur float64 = 0.85
	const usdRub float64 = 82.82

	eurRub := usdRub / usdEur
	fmt.Print(eurRub)
}
