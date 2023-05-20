package Algoritmos_de_busqueda

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	// Caso de prueba: valor objetivo se encuentra en el arreglo
	target := 23
	expectedIndex := 5
	index := binarySearch(arr, target)
	if index != expectedIndex {
		t.Errorf("La búsqueda binaria falló. Se esperaba el índice %d, se obtuvo %d", expectedIndex, index)
	}

	// Caso de prueba: valor objetivo no se encuentra en el arreglo
	target = 30
	expectedIndex = -1
	index = binarySearch(arr, target)
	if index != expectedIndex {
		t.Errorf("La búsqueda binaria falló. Se esperaba el índice %d, se obtuvo %d", expectedIndex, index)
	}
}

func TestBinarySearchIterative(t *testing.T) {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	// Caso de prueba: valor objetivo se encuentra en el arreglo
	target := 23
	expectedIndex := 5
	index := binarySearchIterative(arr, target)
	if index != expectedIndex {
		t.Errorf("La búsqueda binaria iterativa falló. Se esperaba el índice %d, se obtuvo %d", expectedIndex, index)
	}

	// Caso de prueba: valor objetivo no se encuentra en el arreglo
	target = 30
	expectedIndex = -1
	index = binarySearchIterative(arr, target)
	if index != expectedIndex {
		t.Errorf("La búsqueda binaria iterativa falló. Se esperaba el índice %d, se obtuvo %d", expectedIndex, index)
	}
}
