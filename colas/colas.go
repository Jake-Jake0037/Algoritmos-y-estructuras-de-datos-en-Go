package colas

import (
	"errors"
)

// Definición de la estructura de la cola
type Queue struct {
	items []int
}

// Función para crear una nueva cola vacía
func NewQueue() *Queue {
	return &Queue{}
}

// Función para verificar si la cola está vacía
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Función para obtener el tamaño de la cola
func (q *Queue) Size() int {
	return len(q.items)
}

// Función para agregar un elemento al final de la cola
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Función para eliminar y obtener el elemento del frente de la cola
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("la cola está vacía")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Función para obtener el elemento del frente de la cola sin eliminarlo
func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("la cola está vacía")
	}
	return q.items[0], nil
}
