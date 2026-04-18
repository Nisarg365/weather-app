package main

import (
	"fmt"
)

func main() {
	var temp int
	for {
		fmt.Println("--- Smart Weather Bot ---")
		fmt.Print("What is the temperature in Celsius? ")

		_, err := fmt.Scan(&temp)

		if err != nil {
			fmt.Println("❌ Error: Please enter a whole number!")

			var discard string
			fmt.Scanln(&discard)

			continue
		}

		if temp == 999 {
			fmt.Println("Exit")
			break
		} else if temp == 21 {
			fmt.Println("21°C? That is the perfect temperature for a PhD student! 🎓")
		} else if temp > 25 && temp != 999 {
			fmt.Println("Wear a T-shirt ☀️")
		} else if temp < 15 {
			fmt.Println("Bring a Heavy coat ❄️")
		} else {
			fmt.Println("Bring a jacket 1🧥")
		}
	}
}
