package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	fila := []string{}
	lastindex := -1
	firstindex := -1

	if len(list) == 0 {
		return []string{}
	}

	for i := range list {

		if list[i] == first && firstindex == -1 {
			firstindex = i

			if i > 0 {
				fila = append(fila, list[:i]...)
			}
		}

		if list[i] == last && firstindex != -1 && lastindex == -1 {
			lastindex = i //

			if i < len(list)-1 {
				fila = append(fila, list[i+1:]...)
			}
		}
	}

	// 3. Prüfen, ob beide gefunden wurden
	if firstindex == -1 || lastindex == -1 {
		return []string{}
	}

	return fila
}
