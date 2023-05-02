package main

import "fmt"

type livro struct {
	titulo string
	autor  string
	ano    int
}

func main() {
	var x string
	l := livro{"SA", "Tolkein", 1990}
	fmt.Println("Qual livro deseja escolher?")
	fmt.Scan(&x)
	if x == l.titulo {
		fmt.Print("O titulo do livro é", l.titulo)
		fmt.Print("O autor é", l.autor)
		fmt.Print("O ano de publicação é", l.ano)
	}
}
