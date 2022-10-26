package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// func main() {
// 	var ListaDesordenada = []int{15, 3, 8, 6, 18, 1}
// 	var N int = len(ListaDesordenada)
// 	ListaOrdenada := shellSort(ListaDesordenada, N)
// 	fmt.Println(ListaOrdenada)

// 	var ListaDesordenada2 = []int{15, 3, 8, 6, 18, 1}
// 	ListaOrdenada2 := shellSortInv(ListaDesordenada2, N)
// 	fmt.Println(ListaOrdenada2)
// }

func main() {
	n := 1000000 //definir tamaño del arreglo
	arregloDeNumeros := generarNumeros(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //se elige un número al azar para buscar
	var pos int = r.Intn(n-0) + 0                        //se elige un número al azar para buscar
	busquedaNumerica := arregloDeNumeros[pos]            //se elige un número al azar para buscar

	//para arreglo ordenado m-M
	fmt.Println("BUSQUEDA CON ARREGLO ORDENADO DE MENOR A MAYOR")
	arregloOrdenado := shellSort(arregloDeNumeros, len(arregloDeNumeros))
	resultadoBusquedaNumericaBinaria := busquedaBinaria(arregloOrdenado, busquedaNumerica)
	fmt.Printf("Buscando(Binaria) %d en el arreglo... el índice devuelto es %d\n", busquedaNumerica, resultadoBusquedaNumericaBinaria)
	resultadoBusquedaNumericaSecuencial := BusquedaSecuencial(arregloOrdenado, busquedaNumerica)
	fmt.Printf("Buscando(Secuencial) %d en el arreglo... el índice devuelto es %d\n\n", busquedaNumerica, resultadoBusquedaNumericaSecuencial)

	//para arreglo ordenado M-m
	fmt.Println("BUSQUEDA CON ARREGLO ORDENADO DE MAYOR A MENOR")
	arregloOrdenadoMm := shellSortInv(arregloDeNumeros, len(arregloDeNumeros))
	resultadoBusquedaNumericaBinariaMn := busquedaBinaria(arregloOrdenado, busquedaNumerica)
	fmt.Printf("Buscando(Binaria) %d en el arreglo... el índice devuelto es %d\n", busquedaNumerica, resultadoBusquedaNumericaBinariaMn)
	resultadoBusquedaNumericaSecuencialMn := BusquedaSecuencial(arregloOrdenadoMm, busquedaNumerica)
	fmt.Printf("Buscando(Secuencial) %d en el arreglo... el índice devuelto es %d\n\n", busquedaNumerica, resultadoBusquedaNumericaSecuencialMn)

	//para arreglo parcialmente ordenado

	//para arreglo desordenado
	fmt.Println("BUSQUEDA CON ARREGLO DESORDENADO")
	arregloDesordenado := Mezclar(arregloDeNumeros)
	resultadoBusquedaDesordenadaBinaria := busquedaBinaria(arregloOrdenado, busquedaNumerica)
	fmt.Printf("Buscando(Binaria) %d en el arreglo... el índice devuelto es %d\n", busquedaNumerica, resultadoBusquedaDesordenadaBinaria)
	resultadoBusquedaDesordenadaSecuencial := BusquedaSecuencial(arregloDesordenado, busquedaNumerica)
	fmt.Printf("Buscando(Secuencial) %d en el arreglo... el índice devuelto es %d\n\n", busquedaNumerica, resultadoBusquedaDesordenadaSecuencial)
}

func shellSort(A []int, n int) []int {
	var iter = 0
	for interval := int(n / 2); interval > 0; interval /= 2 {
		for j := interval; j < n; j++ {
			for k := j; k >= interval && A[k-interval] > A[k]; k -= interval {
				A[k], A[k-interval] = A[k-interval], A[k]
				iter++
			}
		}
	}
	fmt.Printf("Iteraciones de ordenamiento por Shell: %d \n", iter)
	return A
}

func shellSortInv(A []int, n int) []int {
	for interval := int(n / 2); interval > 0; interval /= 2 {
		for j := interval; j < n; j++ {
			for k := j; k >= interval && A[k-interval] < A[k]; k -= interval {
				A[k], A[k-interval] = A[k-interval], A[k]
			}
		}
	}
	return A
}

func Mezclar(a []int) []int {
	array := make([]int, len(a))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(a))

	for i, v := range perm {
		array[v] = a[i]
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

func BusquedaSecuencial(numeros []int, valor int) int {
	var iter = 0
	// recorrer el arreglo numeros hasta encontrar el valor
	for k, v := range numeros {
		// verificar si el valor iterado coincide con el que se busca
		if v == valor {
			// retornar el indice
			return k
		}
		iter++
	}
	fmt.Printf("Iteraciones de Busqueda Secuencial: %d \n", iter)
	// retornar -1 si ningun valor coincide
	return -1
}

func busquedaBinaria(arreglo []int, busqueda int) (indice int) {
	var iter = 0
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
		iter++
	}
	fmt.Printf("Iteraciones de Busqueda Binaria: %d \n", iter)
	return -1
}
