package mafin

import (
	"fmt"
	"inf1-material/Vorlesungsprojekte/Vorlesung2/fakultaet"
)

func ExampleFactorial() {

	fmt.Println(fakultaet.Factorial(1))
	fmt.Println(fakultaet.Factorial(2))
	fmt.Println(fakultaet.Factorial(3))
	fmt.Println(fakultaet.Factorial(4))
	fmt.Println(fakultaet.Factorial(5))

	// Output:
	//1
	//2
	//6
	//24
	//120
}
func ExampleFactorialV2() {

	fmt.Println(fakultaet.FactorialV2(1))
	fmt.Println(fakultaet.FactorialV2(2))
	fmt.Println(fakultaet.FactorialV2(3))
	fmt.Println(fakultaet.FactorialV2(4))
	fmt.Println(fakultaet.FactorialV2(5))

	// Output:
	//1
	//2
	//6
	//24
	//120

}
