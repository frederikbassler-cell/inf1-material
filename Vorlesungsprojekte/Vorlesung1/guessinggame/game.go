package guessinggame

import "fmt"
import "math/rand/v2"

func GuessingGame() {
	correct := rand.IntN(100)
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess, correct) {
			fmt.Println("Richtig geraten! :-)")
			return
		}
		if guess < correct {
			fmt.Println("Die Zahl ist größer als deine Zahl")
		} else {
			fmt.Println("Die Zahl ist kleiner als deine Zahl")

		}


	}
	fmt.Println("Zu viele falsche Zahlen! :-(")
	fmt.Println("Die richtige Zahl war", correct )

}

