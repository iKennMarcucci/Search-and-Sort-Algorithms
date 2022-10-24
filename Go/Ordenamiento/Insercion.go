package main

import "fmt"

func main() {
	var ListaDesordenada = []int{15, 3, 8, 6, 18, 1}
	ListaOrdenada := Insercion(ListaDesordenada)
	fmt.Println(ListaOrdenada)
}

func Insercion(ListaDesordenada []int) []int {
	var auxiliar int
	for i := 1; i < len(ListaDesordenada); i++ {
		auxiliar = ListaDesordenada[i]
		for j := i - 1; j >= 0 && ListaDesordenada[j] > auxiliar; j-- {
			ListaDesordenada[j+1] = ListaDesordenada[j]
			ListaDesordenada[j] = auxiliar
		}
	}
	return ListaDesordenada
}
