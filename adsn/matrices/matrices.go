package matrices

import "fmt"

//eine matrix ist eine liste an listen

func ExampleMatrices() {

	//matrix definieren (zweidimensional):
	matrix := [][]int{}

	//matrix definieren (n-dimensional):
	//matrixNdim := [][][][][]....[]int{}

	//compiler happy machen
	matrix[0][0] = 1

	matrix_float := [][]float64{

		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	//erste liste ausgeben:
	fmt.Println(matrix_float[0])

	//erstes element ("links oben")
	fmt.Println(matrix_float[0][0])

	//erstes element ("zweite zeile links")
	fmt.Println(matrix_float[1][0])

}
