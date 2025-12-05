package main

import (
	"fmt"
	"inf1-material/Vorlesungsprojekte/Vorlesung1/greet"
)

func main() {

	fmt.Print("Wie heißt du")
	var name string
	fmt.Scanln(&name)
s1 := greet.Greet(name)

fmt.Println(s1)

}