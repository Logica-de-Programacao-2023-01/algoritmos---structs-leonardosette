package main

import "fmt"

type aluno struct {
	nome  string
	idade int
	notas []float64
}

func main() {
	a := aluno{
		nome:  "Gabriel",
		idade: 21,
		notas: []float64{1, 2, 3, 4, 5},
	}
	fmt.Println(media(a))
}
func media(a aluno) float64 {
	var sum float64
	for _, nota := range a.notas {
		sum += nota
	}
	return sum / float64(len(a.notas))
}
