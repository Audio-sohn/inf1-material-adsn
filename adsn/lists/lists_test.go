package lists

import (
	"fmt"
	"unsafe"
)

//TESTS in GO:

// -Dateiname muss in "_test" enden!
// -Funktionen die zu testen sind müssen mit "Example..." im identifier beginnen
// -Output muss spezifiziert werden als kommentar in der Funktionen
// -für performance tests "Benchmark..." am Anfang des identifiers
// -innerhalb der example funktion wird die zu testende funktion mit passenden argumenten aufgerufen

//syntax für die listen in GO

func ExampleLists() {

	// students := []string{"erstes", "zweites"}
	// points := []int{0, 0}

	testArr_float := []float64{}
	testArr_int := []int{}
	testArr_string := []string{}

	var teststring string
	var testint int

	fmt.Print("default array length (float): ")
	fmt.Println(unsafe.Sizeof(testArr_float))

	fmt.Print("default array length (int): ")
	fmt.Println(unsafe.Sizeof(testArr_int))

	fmt.Print("default array length (string): ")
	fmt.Println(unsafe.Sizeof(testArr_string))

	fmt.Print("default data length (int): ")
	fmt.Println(unsafe.Sizeof(testint))

	fmt.Print("default data length (string): ")
	fmt.Println(unsafe.Sizeof(teststring))

	//Output:
	// [erstes zweites]
	// [0 0]

}
