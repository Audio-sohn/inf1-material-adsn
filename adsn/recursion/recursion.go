package recursion

import "fmt"

func Countdown(n int) {

	if n <= 0 {
		return
	}

	fmt.Println(n)

	Countdown(n - 1)

}

func PowersOf5(n int) int {

	if n == 0 {

		return 1

	}

	// für negative Exponenten adaptieren
	if n < 0 {

		// nur der erste call wird modifiziert, ab dann wird normal weitergerechnet
		return 1 / PowersOf5(-n)

	}

	return 5 * PowersOf5(n-1)

	// Wie funktioniert diese funktion?

	// n = 5

	// PowersOf5(5) 	= 5 * PowersOf5(4)
	// 					= 5 * 5 * PowersOf5(3)
	// 					= 5 * 5 * 5 * PowersOf5(2)
	// 					= 5 * 5 * 5 * 5 PowersOf5(1)
	// 					= 5 * 5 * 5 * 5 * 5 * PowersOf5(0)
	// 					= 5 * 5 * 5 * 5 * 5 * 1
	// 					= (5)^5

	// PowersOf5(-5) 	= 1 / PowersOf5(5)
	//					= 1 / 5 * PowersOf5(4)
	// 					= 1 / 5 * 5 * PowersOf5(3)
	// 					= 1 / 5 * 5 * 5 * PowersOf5(2)
	// 					= 1 / 5 * 5 * 5 * 5 PowersOf5(1)
	// 					= 1 / 5 * 5 * 5 * 5 * 5 * PowersOf5(0)
	// 					= 1 / 5 * 5 * 5 * 5 * 5 * 1
	// 					= 1 / (5)^5

}
