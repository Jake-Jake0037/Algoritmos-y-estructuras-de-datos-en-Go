package ordenamento

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 3, 8, 2, 1, 4}
	expectedResult := []int{1, 2, 3, 4, 5, 8, 3} //Hice que no pasara el test a proposito.

	sortedArr := QuickSort(arr)

	if !reflect.DeepEqual(sortedArr, expectedResult) {
		t.Errorf("El arreglo ordenado es incorrecto. Se esperaba %v pero se obtuvo %v", expectedResult, sortedArr)
	}
}
