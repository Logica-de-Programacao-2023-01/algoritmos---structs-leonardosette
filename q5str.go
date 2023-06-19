package main

import "fmt"

//Usando as mesmas structs do exercício anterior, escreva uma função
//que receba um título de música como parâmetro e retorne uma lista
//das Playlists que possuem esse título.

type Musica struct {
	titulo  string
	artista string
	duracao int
}

func encontrarMusicas(playlist []Playlist, nomeMusica string) []Playlist {
	var resultado []Playlist
	for _, playlist := range playlist {
		for _, musica := range playlist.musicas {
			if musica.titulo == nomeMusica {
				resultado = append(resultado, playlist)
			}
		}
	}
	return resultado
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func main() {
	p := Playlist{
		nome: "giovana",
		musicas: []Musica{
			{
				titulo:  "Run",
				artista: "OneRepublic",
				duracao: 3,
			},
			{
				titulo:  "Pumped Up Kicks",
				artista: "Madism, MKJ, Felix Samuel",
				duracao: 3,
			},
		},
	}
	fmt.Println(p)
}
