package main

import (
	"fmt"
	"math"
)

//Crie uma struct chamada Circulo com o campo "raio".
//Escreva uma função que receba um Circulo como parâmetro e
//calcule a área do círculo (área = pi * raio * raio).
//Dica: use a constante math.Pi para representar o número pi.

type Circulo struct {
	raio float64
	area float64
}

func main() {
	p := Circulo{raio: 10}
	fmt.Println(areaCirculo(p))
}

func areaCirculo(p Circulo) (area float64) {
	area = math.Pi * p.raio * p.raio
	return area
}
