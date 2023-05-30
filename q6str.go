package main

import (
	"fmt"
)

func registerUser(names []string) []string {
	database := make(map[string]bool)
	responses := make([]string, len(names))

	for i, name := range names {
		if !database[name] {
			database[name] = true
			responses[i] = "OK"
		} else {
			newName := generateNewName(name, database)
			database[newName] = true
			responses[i] = newName
		}
	}

	return responses
}

func generateNewName(name string, database map[string]bool) string {
	i := 1
	newName := fmt.Sprintf("%s%d", name, i)

	for database[newName] {
		i++
		newName = fmt.Sprintf("%s%d", name, i)
	}

	return newName
}

func main() {
	names := []string{"alex", "steven", "alex", "steven", "alex", "steven", "alex", "steven", "alex", "steven"}

	responses := registerUser(names)

	for _, response := range responses {
		fmt.Println(response)
	}
}
