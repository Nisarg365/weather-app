package main

import (
	"fmt"
)

func main() {
	var temp int

	fmt.Println("--- Smart Weather Bot ---")
	fmt.Print("What is the temperature in Celsius? ")
	fmt.Scan(&temp)
	if temp > 25 {
		fmt.Println("Wear a T-shirt ☀️")
	} else if temp < 15 {
		fmt.Println(" Bring a Heavy coat ❄️")
	} else if temp >= 15 && temp <= 25 {
		fmt.Println("Bring a jacket 🧥")
	} else if temp == 21 {
		fmt.Println("21°C? That is the perfect temperature for a PhD student! 🎓")
	}
}

