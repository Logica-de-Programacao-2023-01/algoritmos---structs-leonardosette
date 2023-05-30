package main

import "fmt"

type retangulo struct {
	altura float64
	base   float64
	area   int
}

func main() {
	var a retangulo
	fmt.Println("Informe a altura")
	fmt.Scan(&a.altura)
	fmt.Println("Informe a base")
	fmt.Scan(&a.base)

	fmt.Println(area(a))
}
func area(a retangulo) float64 {
	return a.altura * a.base
}
