package game

import (
	"fmt"
	"inf1-material/adsn/service"
	"math"
	"math/rand"
)

func GuessingGame() {

	x_random := rand.Intn(10)

	y_guess := service.ReadNumber()

	z := math.Pow(2, float64(x_random))

	for n := 0; n < 2; n++ {

		if x_random == y_guess {
			fmt.Println("Richtig! Gut Gemacht!\r")
			return
		}
		fmt.Println("Falsch, versuchs noch einmal!\r")
		y_guess = service.ReadNumber()
	}

	fmt.Println("Zu oft falsch geraten !!! ---> Verloren!")
	fmt.Print("Das quadrat der gesuchten zahl war: ")
	fmt.Println(z)

}
