package Algoritmos_de_busqueda

import "testing"

func TestLinearSearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	target := 3
	expectedIndex := 2

	index := LinearSearch(array, target)

	if index != expectedIndex {
		t.Errorf("Índice esperado: %d, pero se obtuvo: %d", expectedIndex, index)
	}
}
