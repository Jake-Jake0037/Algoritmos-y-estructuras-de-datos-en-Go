package pilas

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	// Verificar que la pila esté vacía inicialmente
	if !stack.IsEmpty() {
		t.Error("Se esperaba que la pila estuviera vacía")
	}

	// Agregar elementos a la pila
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Verificar el tamaño de la pila
	if stack.Size() != 3 {
		t.Errorf("El tamaño de la pila es incorrecto. Se esperaba 3 pero se obtuvo %d", stack.Size())
	}

	// Verificar el elemento superior de la pila
	topItem, err := stack.Peek()
	if err != nil {
		t.Errorf("Error al obtener el elemento superior de la pila: %s", err.Error())
	} else if topItem != 3 {
		t.Errorf("El elemento superior de la pila es incorrecto. Se esperaba 3 pero se obtuvo %d", topItem)
	}

	// Extraer elementos de la pila
	item1, err := stack.Pop()
	if err != nil {
		t.Errorf("Error al extraer elemento de la pila: %s", err.Error())
	} else if item1 != 3 {
		t.Errorf("El elemento extraído de la pila es incorrecto. Se esperaba 3 pero se obtuvo %d", item1)
	}

	item2, err := stack.Pop()
	if err != nil {
		t.Errorf("Error al extraer elemento de la pila: %s", err.Error())
	} else if item2 != 2 {
		t.Errorf("El elemento extraído de la pila es incorrecto. Se esperaba 2 pero se obtuvo %d", item2)
	}

	// Verificar que la pila esté vacía después de extraer elementos
	if !stack.IsEmpty() {
		t.Error("Se esperaba que la pila estuviera vacía después de extraer elementos")
	}

	// Verificar que el orden de los elementos sea correcto
	expectedResult := []int{1}
	if !reflect.DeepEqual(stack.items, expectedResult) {
		t.Errorf("El orden de los elementos en la pila es incorrecto. Se esperaba %v pero se obtuvo %v", expectedResult, stack.items)
	}
}
