package Algoritmos_de_busqueda

// LinearSearch realiza una búsqueda lineal en un array para encontrar el elemento objetivo.
// Devuelve el índice del elemento objetivo si se encuentra, de lo contrario, devuelve -1.

func LinearSearch(array []int, target int) int {
	// Recorre secuencialmente cada elemento del array
	for i, value := range array {
		// Compara el valor actual con el objetivo
		if value == target {
			// Si se encuentra el objetivo, devuelve el índice correspondiente
			return i
		}
	}
	// Si no se encuentra el objetivo, devuelve -1
	return -1
}
