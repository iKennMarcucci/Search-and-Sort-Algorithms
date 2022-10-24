package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generarNumeros(valor int) []int {
	//n := 15
	//fmt.Println(n)
	//var numeros [15]int
	numeros := make([]int, valor)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ale int
	for f := 0; f < valor; f++ {
		ale = r.Intn(99999-10000) + 10000
		numeros[f] = ale
	}
	fmt.Println(numeros)
	return numeros
}

func BusquedaSecuencial(numeros []int, valor int) int {
	// recorrer el arreglo numeros hasta encontrar el valor
	for k, v := range numeros {
		// verificar si el valor iterado coincide con el que se busca
		if v == valor {
			// retornar el indice
			return k
		}
	}
	// retornar -1 si ningun valor coincide
	return -1
}

func main() {
	numeros := generarNumeros(2)
	//fmt.Println(numeros)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var pos int = r.Intn(10-0) + 0
	buscar := numeros[pos]
	fmt.Println("Número a buscar = ", buscar)
	fmt.Print("Posición = ")
	fmt.Println(BusquedaSecuencial(numeros, buscar))
}
