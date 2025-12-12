package main 
				//i = i + 1 oder i++ oder i += 1 für addieren 
import "fmt"

func fizzbuzz_v1(){
	for i:= 1; i < 42; i++ {
		if i % 5 == 0 && i % 7 == 0 {
			fmt.Println("fizzbuzz")
			continue
		} 
		if i % 5 == 0  {
			fmt.Println("fizz")
			continue
		}
		if i % 7 == 0 {
			fmt.Println("buzz")
			continue
		} 
			fmt.Println(i)
		
	}

}

func fizzbuzz_allgemein(x, y int){

	for i:= 1; i < 42; i++ {
		if i % x == 0 && i % y == 0 {
			fmt.Println("fizzbuzz")
			continue
		} 
		if i % x == 0  {
			fmt.Println("fizz")
			continue
		}
		if i % y == 0 {
			fmt.Println("buzz")
			continue
		} 
			fmt.Println(i)
		
	}


}



func main(){
	//fizzbuzz_v1()




	var x int 
	fmt.Print("Nenne mir ein Fizz: ")
	fmt.Scan(&x) 

	var y int 
	fmt.Print("Nenne mir ein Buzz: ")
	fmt.Scan(&y) 

		fizzbuzz_allgemein(x, y)
}