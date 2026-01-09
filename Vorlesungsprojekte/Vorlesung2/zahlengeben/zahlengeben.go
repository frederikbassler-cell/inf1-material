package main

import "fmt"

func GetInput() int {
	fmt.Println("Gib eine Zahl zwischen 1 und 10 ein: ")
	var n int
	fmt.Scan(&n)

	if n > 0 && n <= 10 {
	return n
	}

	fmt.Println("DU pisser was war falsch!")
	return GetInput()
}




func main() {
n:= GetInput()
fmt.Printf("%d ist eine gute Zahl.", n)


}