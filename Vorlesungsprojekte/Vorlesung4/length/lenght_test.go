package length

import "fmt"

type length int

func (l length) Centimeters() int {
	return int(l)
}

func (l length) Meters() int {
	return l.Centimeters() / 100
}

func (l *length) Scale(factor int){
	*l= length(*l * length(factor))
}

func Example() {

	var a length = 10000
	//var b int = 2

	fmt.Println(a)
	// fmt.Println(a * b) // Nicht erlaubt

	fmt.Println(a.Centimeters())
	fmt.Println(a.Meters())
	

	// Output:

}
