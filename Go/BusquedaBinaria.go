/*
Búsqueda binaria en arreglos de cadenas y números usando Go
@author parzibyte
Visita: parzibyte.me/blog
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	arregloDeNumeros := generarNumeros(10)
	arregloOrdenado := BubbleSort(arregloDeNumeros)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var pos int = r.Intn(10-0) + 0
	busquedaNumerica := arregloDeNumeros[pos]
	resultadoBusquedaNumerica := busquedaBinariaRecursiva(arregloOrdenado, busquedaNumerica, 0, len(arregloOrdenado)-1)
	fmt.Printf("[Recursivo] Buscando %d en %v... el índice devuelto es %d\n", busquedaNumerica, arregloOrdenado, resultadoBusquedaNumerica)
}

func Burbuja(ListaDesordenada []int) []int {
	var auxiliar int
	for i := 0; i < len(ListaDesordenada)-1; i++ {
		for j := 0; j < len(ListaDesordenada); j++ {
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
	//fmt.Println(numeros)
	return numeros
}

func busquedaBinariaRecursiva(arreglo []int, busqueda, izquierda, derecha int) (indice int) {
	if izquierda > derecha {
		return -1
	}
	indiceDelMedio := int(math.Floor((float64(izquierda+derecha) / 2)))
	elementoDelMedio := arreglo[indiceDelMedio]
	if elementoDelMedio == busqueda {
		return indiceDelMedio
	}
	if busqueda < elementoDelMedio {
		derecha = indiceDelMedio - 1
	} else {
		izquierda = indiceDelMedio + 1
	}
	return busquedaBinariaRecursiva(arreglo, busqueda, izquierda, derecha)
}

func busquedaBinaria(arreglo []int, busqueda int) (indice int) {
	izquierda := 0
	derecha := len(arreglo) - 1

	/*
	   Recordemos que For is Go's "while"
	   https://tour.golang.org/flowcontrol/3
	*/
	for izquierda <= derecha {
		indiceDelMedio := int(math.Floor((float64(izquierda+derecha) / 2)))
		elementoDelMedio := arreglo[indiceDelMedio]
		if elementoDelMedio == busqueda {
			return indiceDelMedio
		}
		if busqueda < elementoDelMedio {
			derecha = indiceDelMedio - 1
		} else {
			izquierda = indiceDelMedio + 1
		}
	}
	return -1
}
