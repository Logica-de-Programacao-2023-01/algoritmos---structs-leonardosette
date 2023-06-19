package main

import (
	"fmt"
)

//Crie uma struct chamada Animal com os campos "nome",
//"espécie", "idade" e "som". Escreva funções que
//permitam modificar o som que o animal faz e uma função
//que imprima as informações do animal e o som que ele
//faz.

type Animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func (p Animal) modificarSom(som string) Animal {
	p.som = som
	return p
}
func (p Animal) imprimirInformações() {
	fmt.Println("Nome:", p.nome)
	fmt.Println("Especie:", p.especie)
	fmt.Println("Idade:", p.idade)
	fmt.Println("Som:", p.som)
}

func main() {
	p := Animal{
		nome:    "jeremias",
		especie: "passaro",
		idade:   7,
		som:     "piado",
	}
	p = p.modificarSom("Sem som")
	p.imprimirInformações()
}
