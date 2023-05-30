package main

import (
	"fmt"
)

func canBuyEnoughFood(dogs int, cats int, stock map[string]int) bool {
	dogFood := stock["Ração para cachorro"]
	catFood := stock["Ração para gato"]
	universalFood := stock["Ração universal"]

	if dogs < 0 || cats < 0 {
		return false
	}

	totalDogFood := dogs
	totalCatFood := cats

	if dogFood < totalDogFood {
		return false
	}

	if catFood < totalCatFood {
		return false
	}

	totalFood := totalDogFood + totalCatFood

	if universalFood < totalFood-dogFood-catFood {
		return false
	}

	return true
}

func main() {
	dogs := 2
	cats := 3

	stock := map[string]int{
		"Ração para cachorro": 2,
		"Ração para gato":     2,
		"Ração universal":     2,
	}

	if dogs < 0 || cats < 0 {
		fmt.Println("Números negativos não são permitidos.")
		return
	}

	result := canBuyEnoughFood(dogs, cats, stock)

	fmt.Println("É possível:", result)
}
