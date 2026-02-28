package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// PrefixBelow10 erwartet eine Liste "list" von Zahlen und liefert
// die längste Teil-Liste, mit der "list" beginnt und die nur Zahlen < 10 enthält.
func PrefixBelow10(list []int) []int {
	count := 0
	if len(list) == 0 {
		return []int{}
	}
	pref := []int{}
	for _, v := range list {

		if count == 0 {

			if v == 10 {

				count++

			} else {
				pref = append(pref, v)
			}

		}

	}
	return pref
}
