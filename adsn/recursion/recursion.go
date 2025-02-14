package recursion

import "fmt"

func Countdown(n int) {

	if n <= 0 {

		return

	}

	fmt.Println(n)

	Countdown(n - 1)

}
