package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	result := []int{}
	summe := 0
	longest := []int{}

	if len(l1) > len(l2) {
		longest = l1
	} else {
		longest = l2
	}

for i := 0; i < len(longest); i++ {


		if len(l1) == 0 {
			summe = 0 + l2[i]
		} else if len(l2) == 0 {
			summe = l1[i] + 0

		
		} else if len(l1) < i+1 {
			summe = l2[i] + l1[len(l1)-1]

		} else if len(l2) < i+1 {
			summe = l1[i] + l2[len(l2)-1]

		
		} else {
			summe = l1[i] + l2[i]
		}

		result = append(result, summe)
	}

	return result
}
