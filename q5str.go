package main

import (
	"fmt"
	"math"
)

type Shirt struct {
	Size  string
	Price int
}

func calculatePriceAverage(shirts []Shirt) (float64, error) {
	minPrice := math.MaxInt64
	maxPrice := math.MinInt64

	for _, shirt := range shirts {
		if shirt.Price < minPrice {
			minPrice = shirt.Price
		}
		if shirt.Price > maxPrice {
			maxPrice = shirt.Price
		}
	}

	if minPrice == math.MaxInt64 || maxPrice == math.MinInt64 {
		return 0, fmt.Errorf("Não foi possível calcular a média")
	}

	average := float64(maxPrice+minPrice) / 2
	return average, nil
}

func main() {
	shirts := []Shirt{
		{
			Size:  "M",
			Price: 10,
		},
		{
			Size:  "XXL",
			Price: 20,
		},
		{
			Size:  "S",
			Price: 7,
		},
		{
			Size:  "XXXXXXXS",
			Price: 5,
		},
	}

	average, err := calculatePriceAverage(shirts)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Maior:", maxPrice)
		fmt.Println("Menor:", minPrice)
		fmt.Println("Média:", average)
	}

	shirts = []Shirt{
		{
			Size:  "S",
			Price: 10,
		},
		{
			Size:  "XXL",
			Price: 20,
		},
		{
			Size:  "S",
			Price: 4,
		},
		{
			Size:  "XXXXXXXL",
			Price: 5,
		},
		{
			Size:  "XXXXXXXL",
			Price: 25,
		},
	}

	average, err = calculatePriceAverage(shirts)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Maior:", maxPrice)
		fmt.Println("Menor:", minPrice)
		fmt.Println("Média:", average)
	}
}
