package main

import "fmt"

func main() {
	var ListaDesordenada = []int{15, 3, 8, 6, 18, 1}
	var N int = len(ListaDesordenada)
	ListaOrdenada := quicksort(ListaDesordenada, 0, N-1)
	fmt.Println(ListaOrdenada)
}

func quicksort(ListaDesordenada []int, izq int, der int) []int {
	pivote := ListaDesordenada[izq] // tomamos primer elemento como pivote
	i := izq                        // i realiza la búsqueda de izquierda a derecha
	j := der                        // j realiza la búsqueda de derecha a izquierda
	var aux int

	for i < j { // mientras no se crucen las búsquedas
		for ListaDesordenada[i] <= pivote && i < j {
			i++
		} // busca elemento mayor que pivote
		for ListaDesordenada[j] > pivote {
			j--
		} // busca elemento menor que pivote
		if i < j { // si no se han cruzado
			aux = ListaDesordenada[i] // los intercambia
			ListaDesordenada[i] = ListaDesordenada[j]
			ListaDesordenada[j] = aux
		}
	}
	ListaDesordenada[izq] = ListaDesordenada[j] // se coloca el pivote en su lugar de forma que tendremos
	ListaDesordenada[j] = pivote                // los menores a su izquierda y los mayores a su derecha
	if izq < j-1 {
		quicksort(ListaDesordenada, izq, j-1) // ordenamos subarray izquierdo
	}

	if j+1 < der {
		quicksort(ListaDesordenada, j+1, der) // ordenamos subarray derecho
	}
	return ListaDesordenada
}
