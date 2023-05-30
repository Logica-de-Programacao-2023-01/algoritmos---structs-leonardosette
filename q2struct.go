package main

import "fmt"

type Participant struct {
	Name string
	Role string
}

func calculateMaxTeams(participants []Participant) int {
	programmers := 0
	mathematicians := 0

	for _, participant := range participants {
		if participant.Role == "Programmer" {
			programmers++
		} else if participant.Role == "Mathematician" {
			mathematicians++
		}
	}
	
	maxTeams := 0
	if programmers >= 1 && mathematicians >= 1 {
		minTeamMembers := min(programmers/4, mathematicians/4)
		maxTeams = minTeamMembers
	}

	return maxTeams
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

	maxTeams := calculateMaxTeams(participants)

	fmt.Println("Número máximo de equipes:", maxTeams)
}
