package ordenamento

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 3, 8, 2, 1, 4}
	expectedResult := []int{1, 2, 3, 4, 5, 8}

	sortedArr := MergeSort(arr)

	if !reflect.DeepEqual(sortedArr, expectedResult) {
		t.Errorf("El arreglo ordenado es incorrecto. Se esperaba %v pero se obtuvo %v", expectedResult, sortedArr)
	}
}
