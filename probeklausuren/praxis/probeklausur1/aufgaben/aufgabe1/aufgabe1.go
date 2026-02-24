package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ShortestAbc.
MAX. PUNKTE: 10
*/

// ShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
//
// Hinweis: Die Funktion muss nur mit kurzen Strings der Länge < 100 funktionieren.
func ShortestAbc(list []string) string {
	shortest := ""

	if len(list) == 0 {
		return ""
	}

	for _, v := range list {
		if len(v) >= 3 {
			if v[0] == 'a' && v[1] == 'b' && v[2] == 'c' {
				
				if len(v) < len(shortest) || shortest == ""{
					shortest = v 
				}
			}
		}
	}

	return shortest
}
