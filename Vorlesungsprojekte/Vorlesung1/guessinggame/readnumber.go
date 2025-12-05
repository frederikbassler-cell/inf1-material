package guessinggame

import "fmt"

func ReadNumber() int {

	var n int 
	fmt.Print("Rate eine Zahl!")
	fmt.Scan(&n) //cin aus cpp
	return n

}