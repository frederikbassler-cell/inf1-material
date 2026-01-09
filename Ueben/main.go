package main

import (
	"fmt"
	"inf1-material/Ueben/calculator"
)

func main() {

	fmt.Println("23 + 42 =", calculator.Add(23, 42))
	fmt.Println(calculator.Divide(17, 3))
	fmt.Println(calculator.Sum(1, 10))
	fmt.Println(calculator.SumUntil(1000))
	fmt.Println(calculator.IsSquareNumber(16))
	calculator.DemoCollections()
}
