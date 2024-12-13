package service

import (
	"fmt"
)

func ReadNumber() int {

	var n int
	fmt.Print("Rate eine Zahl zwischen 1 - 10: ")
	fmt.Scan(&n)
	return n

}
