package main

import (
	"fmt"
)

type Participant struct {
	Name string
	Role string
}

func calculateMaxTeams(participants []Participant) (int, error) {
	programmers := 0
	mathematicians := 0
	
	for _, participant := range participants {
		if participant.Role == "Programmer" {
			programmers++
		} else if participant.Role == "Mathematician" {
			mathematicians++
		}
	}

	for _, participant := range participants {
		if participant.Name == "Pedro" || participant.Name == "Vanessa" {
			if programmers < 1 {
				return 0, fmt.Errorf("Não há programadores suficientes")
			}
		} else if participant.Name == "Tônia" || participant.Name == "João" {
			if mathematicians < 1 {
				return 0, fmt.Errorf("Não há matemáticos suficientes")
			}
		}
	}

	maxTeams := 0
	if programmers >= 1 && mathematicians >= 1 {
		minTeamMembers := min(programmers/4, mathematicians/4)
		maxTeams = minTeamMembers
	}

	return maxTeams, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	participants := []Participant{
		{
			Name: "Pedro",
			Role: "Programmer",
		},
		{
			Name: "Vanessa",
			Role: "Programmer",
		},
		{
			Name: "Tônia",
			Role: "Mathematician",
		},
		{
			Name: "João",
			Role: "Mathematician",
		},
	}

	maxTeams, err := calculateMaxTeams(participants)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Número máximo de equipes:", maxTeams)
	}
}
