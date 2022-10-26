package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	n := 100 //definir tamaño del arreglo 100, 1000, 1000000, 1000000000
	arregloDeNumeros := generarNumeros(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //se elige un número al azar para buscar
	var pos int = r.Intn(n-0) + 0                        //se elige un número al azar para buscar
	busquedaNumerica := arregloDeNumeros[pos]            //se elige un número al azar para buscar

	//para arreglo ordenado m-M
	fmt.Println("BUSQUEDA CON ARREGLO ORDENADO DE MENOR A MAYOR")
	sort.Ints(arregloDeNumeros)
	resultadoBusquedaNativa := sort.SearchInts(arregloDeNumeros, busquedaNumerica)
	fmt.Printf("Buscando %d en el arreglo... el índice devuelto es %d\n\n", busquedaNumerica, resultadoBusquedaNativa)
	//fmt.Println(arregloDeNumeros)

	//para arreglo ordenado M-m
	fmt.Println("BUSQUEDA CON ARREGLO ORDENADO DE MAYOR A MENOR")
	sort.Sort(sort.Reverse(sort.IntSlice(arregloDeNumeros)))
	resultadoBusquedaNativaMn := sort.SearchInts(arregloDeNumeros, busquedaNumerica)
	fmt.Printf("Buscando %d en el arreglo... el índice devuelto es %d\n\n", busquedaNumerica, resultadoBusquedaNativaMn)
	//fmt.Println(arregloDeNumeros)

	//para arreglo parcialmente ordenado

	//para arreglo desordenado
	fmt.Println("BUSQUEDA CON ARREGLO DESORDENADO")
	arregloDesordenado := Mezclar(arregloDeNumeros)
	resultadoBusquedaDesordenadaNativa := sort.SearchInts(arregloDesordenado, busquedaNumerica)
	fmt.Printf("Buscando %d en el arreglo... el índice devuelto es %d\n\n", busquedaNumerica, resultadoBusquedaDesordenadaNativa)
	//fmt.Println(arregloDesordenado)
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
