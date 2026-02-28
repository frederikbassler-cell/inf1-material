package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// FilterDigits erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Ziffern entfernt werden.
// Alle anderen Zeichen sollen unverändert bleiben.
func FilterDigits(s string) string {
	filtered := []rune{}

	if len(s) == 0{
		return ""
	}
	for _, v := range s {
		if v == '1' || v == '2' || v == '3' || v == '4' || v == '5' || v == '6' || v == '7' || v == '8' || v == '9' || v == '0' {
			
		} else {
			filtered = append(filtered, v )
		}
	}

	return string(filtered)
}
