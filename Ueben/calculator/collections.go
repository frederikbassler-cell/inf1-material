package calculator

import "fmt"


func DemoCollections() {

	primesArray := [5]int{2, 3, 5, 7, 11}

	fmt.Println(primesArray)

	primesSlice := make([]int, 5, 7)

	primesSlice = append(primesSlice, 2)
	primesSlice = append(primesSlice, 3)
	primesSlice = append(primesSlice, 5)
	fmt.Println(primesSlice)
	fmt.Println(len(primesSlice)) // Länge
	fmt.Println(cap(primesSlice)) //Kapazität

	for index, value := range primesSlice {		// index = welche Position in der liste und value welche Zahl/Wert dort ist. 
		fmt.Println(index, value)

	}

}
