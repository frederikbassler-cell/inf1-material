package main

import (
	"fmt"
	numbers "inf1-material/Aufgaben/Numbers/count"
)

func main() {
	result := numbers.CountDivisors()
	fmt.Println("Anzahl der Teiler:", result)
}
