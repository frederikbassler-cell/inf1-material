package numbers

import "fmt"

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors() int {
	var n int
	fmt.Print("Nenne mir die Zahl, von der du die Teiler wissen möchtest: ")
	fmt.Scan(&n) //cin aus cpp
	
	if n < 1 {
		fmt.Println("Bitte eine ganze Zahl >= 1 eingeben.")
		return 0
	}

	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				count += 1
			} else {
				count += 2
			}
		}
	}

	return count
}
