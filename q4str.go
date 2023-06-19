package main

import "fmt"

//Crie uma struct chamada Playlist com os campos "nome" e "músicas".
//O campo "músicas" deve ser um slice de structs, cada uma
//representando uma música com os campos "título", "artista"
//e "duração". Escreva uma função que receba uma Playlist como
//parâmetro e imprima o nome da playlist, o tempo total da
//playlist e as informações de cada música.

type musicas struct {
	titulo  string
	artista string
	duracao float64
}

type Playlists struct {
	nome   string
	musica []musicas
}

func main() {
	p := Playlists{
		nome: "pop",
		musica: []musicas{
			{
				titulo:  "Run",
				artista: "OneRepublic",
				duracao: 2.48,
			},
			{
				titulo:  "Pumped Up Kicks",
				artista: "Madism, MKJ, Felix Samuel",
				duracao: 2.25,
			},
		},
	}
	imprimaPlaylist(p)

}

func imprimaPlaylist(p Playlists) {
	fmt.Println("Nome da Playlist:", p.nome)
	var soma float64 = 0
	for _, a := range p.musica {
		soma += a.duracao
		fmt.Println("Duracao total:", soma)
		fmt.Println("Titulo:", a.titulo)
		fmt.Println("Artista:", a.artista)
		fmt.Println("Duracao:", a.duracao)
		fmt.Println("")
	}

}
