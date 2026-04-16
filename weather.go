package main

import (
	"fmt"
)

func main() {
	var temp int

	fmt.Println("--- Smart Weather Bot ---")
	fmt.Print("What is the temperature in Celsius? ")
	fmt.Scan(&temp)
	// 1. Get the user input here using fmt.Scan
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

//

// 2. Write the if / else if / else logic below
// Above 25: Wear a T-shirt ☀️
// 15 to 25: Bring a jacket 🧥
// Below 15: Heavy coat ❄️
