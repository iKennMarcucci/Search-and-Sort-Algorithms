package main

import "fmt"

func main() {
	var ListaDesordenada = []int{15, 3, 8, 6, 18, 1}
	ListaOrdenada := Burbuja(ListaDesordenada)
	ListaOrdenada2 := BubbleSort(ListaDesordenada)
	fmt.Println(ListaOrdenada)
	fmt.Println(ListaOrdenada2)
}

func Burbuja(ListaDesordenada []int) []int {
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
		for j := 1; j < len(ListaDesordenada); j++ {
			if ListaDesordenada[i] > ListaDesordenada[j] {
				auxiliar = ListaDesordenada[i]
				ListaDesordenada[i] = ListaDesordenada[j]
				ListaDesordenada[j] = auxiliar
			}
		}
	}
	return ListaDesordenada
}

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
