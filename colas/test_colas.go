package colas

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()

	// Verificar que la cola esté vacía inicialmente
	if !queue.IsEmpty() {
		t.Error("Se esperaba que la cola estuviera vacía")
	}

	// Agregar elementos a la cola
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Verificar el tamaño de la cola
	if queue.Size() != 3 {
		t.Errorf("El tamaño de la cola es incorrecto. Se esperaba 3 pero se obtuvo %d", queue.Size())
	}

	// Verificar el elemento del frente de la cola
	frontItem, err := queue.Front()
	if err != nil {
		t.Errorf("Error al obtener el elemento del frente de la cola: %s", err.Error())
	} else if frontItem != 1 {
		t.Errorf("El elemento del frente de la cola es incorrecto. Se esperaba 1 pero se obtuvo %d", frontItem)
	}

	// Eliminar elementos de la cola
	item1, err := queue.Dequeue()
	if err != nil {
		t.Errorf("Error al eliminar elemento de la cola: %s", err.Error())
	} else if item1 != 1 {
		t.Errorf("El elemento eliminado de la cola es incorrecto. Se esperaba 1 pero se obtuvo %d", item1)
	}

	item2, err := queue.Dequeue()
	if err != nil {
		t.Errorf("Error al eliminar elemento de la cola: %s", err.Error())
	} else if item2 != 2 {
		t.Errorf("El elemento eliminado de la cola es incorrecto. Se esperaba 2 pero se obtuvo %d", item2)
	}

	// Verificar que la cola esté vacía después de eliminar elementos
	if !queue.IsEmpty() {
		t.Error("Se esperaba que la cola estuviera vacía después de eliminar elementos")
	}

	// Verificar que el orden de los elementos sea correcto
	expectedResult := []int{3}
	if !reflect.DeepEqual(queue.items, expectedResult) {
		t.Errorf("El orden de los elementos en la cola es incorrecto. Se esperaba %v pero se obtuvo %v", expectedResult, queue.items)
	}
}
