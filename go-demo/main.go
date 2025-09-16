package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.93
	var userKg float64 = 93
	imt := userKg / math.Pow(userHeight, 2)
	fmt.Print(imt)
}
