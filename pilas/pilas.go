package pilas

import (
	"errors"
)

// Definición de la estructura de la pila
type Stack struct {
	items []int
}

// Función para crear una nueva pila vacía
func NewStack() *Stack {
	return &Stack{}
}

// Función para verificar si la pila está vacía
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Función para obtener el tamaño de la pila
func (s *Stack) Size() int {
	return len(s.items)
}

// Función para agregar un elemento a la pila (push)
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Función para eliminar y obtener el elemento superior de la pila (pop)
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("la pila está vacía")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Función para obtener el elemento superior de la pila sin eliminarlo (peek)
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("la pila está vacía")
	}
	return s.items[len(s.items)-1], nil
}
