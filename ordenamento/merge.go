package ordenamento

// Función principal que ordena un arreglo utilizando el algoritmo Merge Sort
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Dividir el arreglo en dos mitades
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	// Ordenar recursivamente las dos mitades
	left = MergeSort(left)
	right = MergeSort(right)

	// Combinar las dos mitades ordenadas
	return merge(left, right)
}

// Función auxiliar para combinar dos arreglos ordenados en uno solo
func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0 // Índice para recorrer el arreglo 'left'
	j := 0 // Índice para recorrer el arreglo 'right'
	k := 0 // Índice para recorrer el arreglo 'result'

	// Combinar los elementos de 'left' y 'right' en orden ascendente
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// Agregar los elementos restantes de 'left', si los hay
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	// Agregar los elementos restantes de 'right', si los hay
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}
