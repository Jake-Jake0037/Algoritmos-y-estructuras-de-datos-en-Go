package busqueda

// La función de búsqueda binaria recursiva toma un arreglo ordenado y un valor objetivo
// y devuelve el índice del valor objetivo si se encuentra en el arreglo,
// o -1 si no se encuentra.
func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchRecursive(arr, target, 0, len(arr)-1)
}

// Función auxiliar recursiva para realizar la búsqueda binaria
func binarySearchRecursive(arr []int, target int, low int, high int) int {
	if low > high {
		return -1 // El valor objetivo no se encuentra en el arreglo
	}

	mid := (low + high) / 2 // Índice central del rango de búsqueda

	if arr[mid] == target {
		return mid // El valor objetivo se encuentra en el índice medio
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, high) // El valor objetivo es mayor, buscar en la mitad derecha
	} else {
		return binarySearchRecursive(arr, target, low, mid-1) // El valor objetivo es menor, buscar en la mitad izquierda
	}
}


// La función de búsqueda binaria iterativa toma un arreglo ordenado y un valor objetivo
// y devuelve el índice del valor objetivo si se encuentra en el arreglo,
// o -1 si no se encuentra.
func BinarySearchIterative(arr []int, target int) int {
	low := 0             // Índice inferior del rango de búsqueda
	high := len(arr) - 1 // Índice superior del rango de búsqueda

	for low <= high {
		mid := (low + high) / 2 // Índice central del rango de búsqueda

		if arr[mid] == target {
			return mid // El valor objetivo se encuentra en el índice medio
		} else if arr[mid] < target {
			low = mid + 1 // El valor objetivo es mayor, buscar en la mitad derecha
		} else {
			high = mid - 1 // El valor objetivo es menor, buscar en la mitad izquierda
		}
	}

	return -1 // El valor objetivo no se encuentra en el arreglo
}
