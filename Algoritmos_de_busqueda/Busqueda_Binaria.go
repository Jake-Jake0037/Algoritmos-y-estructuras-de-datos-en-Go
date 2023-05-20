package Algoritmos_de_busqueda

// La función de búsqueda binaria toma un arreglo ordenado y un valor objetivo
// y devuelve el índice del valor objetivo si se encuentra en el arreglo,
// o -1 si no se encuentra.
func binarySearch(arr []int, target int) int {
	low := 0         // Índice inferior del rango de búsqueda
	high := len(arr) // Índice superior del rango de búsqueda

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

// La función de búsqueda binaria iterativa toma un arreglo ordenado y un valor objetivo
// y devuelve el índice del valor objetivo si se encuentra en el arreglo,
// o -1 si no se encuentra.
func binarySearchIterative(arr []int, target int) int {
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
