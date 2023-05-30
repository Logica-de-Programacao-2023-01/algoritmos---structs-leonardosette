package main

import (
	"fmt"
)

type Video struct {
	ID            int
	Duration      int
	Entertainment int
}

func chooseVideo(t int, videos []Video) (Video, error) {
	var chosenVideo Video
	maxEntertainment := 0

	for _, video := range videos {
		if video.Duration <= t {
			if video.Entertainment > maxEntertainment {
				maxEntertainment = video.Entertainment
				chosenVideo = video
			}
		}
	}

	if maxEntertainment == 0 {
		return Video{}, fmt.Errorf("Não há vídeos adequados dentro do tempo disponível")
	}

	return chosenVideo, nil
}

func main() {
	t := 5
	videos := []Video{
		{
			ID:            1,
			Duration:      2,
			Entertainment: 3,
		},
		{
			ID:            2,
			Duration:      3,
			Entertainment: 4,
		},
		{
			ID:            3,
			Duration:      1,
			Entertainment: 2,
		},
	}

	chosenVideo, err := chooseVideo(t, videos)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Vídeo escolhido: %+v\n", chosenVideo)
	}

	t = 10
	videos = []Video{
		{
			ID:            1,
			Duration:      6,
			Entertainment: 8,
		},
		{
			ID:            2,
			Duration:      7,
			Entertainment: 5,
		},
		{
			ID:            3,
			Duration:      9,
			Entertainment: 9,
		},
		{
			ID:            4,
			Duration:      4,
			Entertainment: 6,
		},
	}

	chosenVideo, err = chooseVideo(t, videos)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Vídeo escolhido: %+v\n", chosenVideo)
	}
}
