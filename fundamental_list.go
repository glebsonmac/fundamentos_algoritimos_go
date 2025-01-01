package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertAtBeginning(data interface{}) {
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) InsertAtEnd(data interface{}) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) Delete(data interface{}) bool {
	if list.head == nil {
		return false
	}

	if list.head.data == data {
		list.head = list.head.next
		return true
	}

	current := list.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) Search(data interface{}) bool {
	current := list.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

// Insere um elemento em uma posição aleatória da lista
func (list *LinkedList) InsertAtRandom(data interface{}, size int) {
	// Gera um número aleatório entre 0 e o tamanho da lista
	rand.Seed(time.Now().UnixNano()) // Inicializa a semente do gerador de números aleatórios
	randomIndex := rand.Intn(size + 1)

	// Cria um novo nó
	newNode := &Node{data: data}

	if randomIndex == 0 {
		// Insere no início
		newNode.next = list.head
		list.head = newNode
		return
	}

	// Encontra o nó anterior à posição de inserção
	current := list.head
	for i := 0; i < randomIndex-1; i++ {
		current = current.next
	}

	// Insere o novo nó
	newNode.next = current.next
	current.next = newNode
}

func main() {
	list := LinkedList{}
	size := 100 // Número de elementos a serem inseridos

	// Insere elementos aleatoriamente
	for i := 0; i < size; i++ {
		list.InsertAtRandom(rand.Intn(10000), i) // Insere um número aleatório entre 0 e 9999
	}

	// Busca pelo primeiro número divisível por 3
	current := list.head
	for current != nil {
		if current.data.(int)%3 == 0 {
			fmt.Println("O primeiro número divisível por 3 é:", current.data)
			break
		}
		current = current.next
	}
}
