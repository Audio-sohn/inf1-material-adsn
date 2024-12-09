package main

import (
	"fmt"
	"math/rand"
)

func ReadNumber() int {

	var n int
	fmt.Print("Rate eine Zahl: ")
	fmt.Scan(&n)
	return n

}

func NumberGood(guess int, number int) bool {

	number = 42
	return guess == number

}

func guessingGame() {

	x := rand.Int()

	y := ReadNumber()

	for n := 1; n < 3; n++ {

		if NumberGood(y, x) {
			fmt.Println("Richtig! Gut Gemacht!\n")
			return
		}
		fmt.Println("Falsch, versuchs noch einmal!\n")
		y = ReadNumber()
	}

	fmt.Println("Zu oft falsch geraten !!! ---> Verloren!")

}

func main() {

	guessingGame()

}
