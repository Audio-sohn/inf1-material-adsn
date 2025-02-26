package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {
	result := []int{}

	equal := false

	// l1 elemente durchiterieren
	for _, item1 := range l1 {

		// flag reset
		equal = false

		// pro l1 element l2 gegenchecken
		for _, item2 := range l2 {

			if item2 == item1 {

				// flag setzen
				equal = true

			}
		}
		// wenn flag gesetzt, appenden überspringen
		if equal {
			continue
		}
		result = append(result, item1)

	}

	// das ganze nochmal für l2
	for _, item2 := range l2 {

		//flag reset
		equal = false

		// pro l2 element l1 gegenchecken
		for _, item1 := range l1 {

			if item2 == item1 {

				// flag setzen
				equal = true

			}
		}
		// wenn flag gesetzt, appenden überspringen
		if equal {
			continue
		}
		result = append(result, item2)

	}

	return result
}
