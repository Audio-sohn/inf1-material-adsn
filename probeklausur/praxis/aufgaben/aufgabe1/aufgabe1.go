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

	found := false
	minlen := 0
	minindex := 0

	for i, speci := range list {

		if len(speci) < 3 {

			continue

		}

		if speci[:3] == "abc" {

			if !found || len(speci) < minlen {

				minlen = len(speci)
				minindex = i
				found = true
				continue
			}

		}

	}

	if found {

		return list[minindex]

	}

	return ""

}
