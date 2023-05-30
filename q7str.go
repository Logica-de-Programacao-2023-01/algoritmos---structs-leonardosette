package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func inverterLista(lista *LinkedList) {
	var prev *Node
	current := lista.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	lista.Head = prev
}

func main() {
	lista := LinkedList{
		Head: &Node{
			Value: 5,
			Next: &Node{
				Value: 20,
				Next: &Node{
					Value: 32,
					Next:  nil,
				},
			},
		},
	}

	fmt.Println("Lista original:")
	printLista(&lista)

	inverterLista(&lista)

	fmt.Println("Lista invertida:")
	printLista(&lista)
}

func printLista(lista *LinkedList) {
	current := lista.Head

	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
