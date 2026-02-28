package aufgabe3

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */
import "math"

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.
func CountSquares(list []int) int {

	if len(list) == 0 {
		return 0
	}

	if list[0] == 1 {
		return CountSquares(list[1:]) + 1
	}

	s := int(math.Sqrt(float64(list[0])))

	if s*s == list[0] {
		return CountSquares(list[1:]) + 1
	} 

	return CountSquares(list[1:])

}
