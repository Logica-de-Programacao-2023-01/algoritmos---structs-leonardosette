package main

import "fmt"

// Crie uma struct chamada Pessoa com os campos "nome", "idade" e
// "endereço". O campo "endereço" deve ser uma outra struct com
//os campos "rua", "número", "cidade" e "estado". Escreva uma
//função que receba uma Pessoa como parâmetro e imprima seu
//endereço completo.

type Endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}

func main() {
	p := Pessoa{nome: "giovana",
		idade: 18,
		endereco: Endereco{
			rua:    "Alvorada",
			numero: 106,
			cidade: "Brasilia",
			estado: "DF"},
	}
	imprimaPessoa(p)
}

func imprimaPessoa(p Pessoa) {
	fmt.Println("Nome:", p.nome)
	fmt.Println("Idade:", p.idade)
	fmt.Println("Rua:", p.endereco.rua)
	fmt.Println("Numero:", p.endereco.numero)
	fmt.Println("Cidade:", p.endereco.cidade)
	fmt.Println("Estado:", p.endereco.estado)
}
