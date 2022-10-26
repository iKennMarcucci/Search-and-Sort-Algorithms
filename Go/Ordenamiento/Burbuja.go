package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// func main() {
// 	var ListaDesordenada = []int{15, 3, 8, 6, 18, 1}
// 	var ListaDesordenada2 = []int{15, 3, 8, 6, 18, 1}
// 	ListaOrdenada := BurbujaInv(ListaDesordenada)
// 	ListaOrdenada2 := BubbleSort(ListaDesordenada2)
// 	fmt.Println(ListaOrdenada)
// 	fmt.Println(ListaOrdenada2)
// }

func main() {
	n := 1000 //definir tamaño del arreglo 100, 1000, 1000000, 1000000000
	arregloDeNumeros := generarNumeros(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //se elige un número al azar para buscar
	var pos int = r.Intn(n-0) + 0                        //se elige un número al azar para buscar
	busquedaNumerica := arregloDeNumeros[pos]            //se elige un número al azar para buscar

	//para arreglo ordenado m-M
	fmt.Println("BUSQUEDA CON ARREGLO ORDENADO DE MENOR A MAYOR")
	arregloOrdenado := BubbleSort(arregloDeNumeros)
	resultadoBusquedaNumericaBinaria := busquedaBinaria(arregloOrdenado, busquedaNumerica)
	fmt.Printf("Buscando(Binaria) %d en el arreglo... el índice devuelto es %d\n", busquedaNumerica, resultadoBusquedaNumericaBinaria)
	resultadoBusquedaNumericaSecuencial := BusquedaSecuencial(arregloOrdenado, busquedaNumerica)
	fmt.Printf("Buscando(Secuencial) %d en el arreglo... el índice devuelto es %d\n\n", busquedaNumerica, resultadoBusquedaNumericaSecuencial)

	//para arreglo ordenado M-m
	fmt.Println("BUSQUEDA CON ARREGLO ORDENADO DE MAYOR A MENOR")
	arregloOrdenadoMm := BurbujaInv(arregloDeNumeros)
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

func BurbujaInv(ListaDesordenada []int) []int {
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
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
	var iter = 0
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
			iter++
		}
	}
	fmt.Printf("Iteraciones de ordenamiento burbuja: %d \n", iter)
	return array
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
