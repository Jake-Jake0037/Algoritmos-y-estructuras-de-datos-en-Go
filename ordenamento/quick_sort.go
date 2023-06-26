package ordenamento

// Función principal que ordena un arreglo utilizando el algoritmo Quicksort
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Selección del pivote: último elemento del arreglo
	pivot := arr[len(arr)-1]

	// Creación de subarreglos izquierdo y derecho
	left := []int{}
	right := []int{}

	// Dividir el arreglo en subarreglos basados en el pivote
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	// Ordenar recursivamente los subarreglos izquierdo y derecho
	sortedLeft := QuickSort(left)
	sortedRight := QuickSort(right)

	// Combinar los subarreglos ordenados junto con el pivote
	return append(append(sortedLeft, pivot), sortedRight...)
}
