package aufgabe5

/*
/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
die auf dem darunter definierten Datentyp arbeitet.
MAX. PUNKTE: 10
*/

// ReplaceEn sucht in der Liste nach dem Eintrag für
// de und ersetzt dessen En-Wert mit dem gegebenen en.
// Gibt es mehrere solche Einträge, soll nur der erste ersetzt werden.
func ReplaceEn(dict []Entry, de, en string) {
	count := 0
	if len(dict) == 0 {
		return
	}
	for i := range dict {
		if count == 0 {
			if dict[i].De == de {
				dict[i].En = en
				count++
			}
		}

	}

	return
}

// Entry repräsentiert einen Eintrag in einem Wörterbuch
type Entry struct {
	De string
	En string
}
