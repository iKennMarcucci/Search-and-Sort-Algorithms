import random

numeros = []


# GENERADOR DE ARRAYS DESORDENADOS
def generarArrayDesordenado():
    for i in range(100):  # 10^2 (Si desea cambiar tamaño, agregar 0's)
        num = random.randint(10000, 99999)
        numeros.append(num)


# GENERADOR DE ARRAYS ORDENADOS MENOS A MAYOR
def generarArrayOrdenado_mM():
    i = 10000
    while i != 10100:  # 10^2 (Si desea cambiar tamaño, agregar 0's)
        num = i
        i += 1
        numeros.append(num)


# GENERADOR DE ARRAYS ORDENADOS MAYOR A MENOR
def generarArrayOrdenado_Mm():
    i = 10100
    while i != 10000:  # 10^2 (Si desea cambiar tamaño, agregar 0's)
        num = i
        i -= 1
        numeros.append(num)


# GENERADOR DE ARRAYS PARCIALMENTE ORDENADOS
def generarArrayParcial():
    i = 10000
    j = 10050
    while i != 10050:  # 10^2 (Si desea cambiar tamaño, agregar 0's)
        while j != 10100:  # 10^2 (Si desea cambiar tamaño, agregar 0's)
            num = j
            j += 1
            numeros.append(num)
        num = i
        i += 1
        numeros.append(num)


# ALGORITMOS DE ORDENAMIENTO POR QUICKSORT
def partition(array, start, end):
    pivot = array[start]
    low = start + 1
    high = end
    while True:
        while low <= high and array[high] >= pivot:
            high = high - 1
        while low <= high and array[low] <= pivot:
            low = low + 1
        if low <= high:
            array[low], array[high] = array[high], array[low]
        else:
            break
    array[start], array[high] = array[high], array[start]
    return high


def Quicksort(array, start, end):
    if start >= end:
        return
    p = partition(array, start, end)
    Quicksort(array, start, p - 1)
    Quicksort(array, p + 1, end)

# ALGORITMO DE BUSQUEDA LINEAL


def busquedaLineal(datos, buscar):
    i = 0
    for x in datos:
        if x == buscar:
            print('BUSQUEDA LINEAL: Número encontrado en la posición',
                  i, 'Iterando', i + 1, 'veces')
            break
        i += 1
    if i == len(datos):
        print('BUSQUEDA LINEAL: No se encontró el número dentro del array.')


# ALGORITMO DE BUSQUEDA BINARIA
def busquedaBinaria(datos, buscar):
    izq = 0
    der = len(datos) - 1
    iteraciones = 0
    while izq <= der:
        iteraciones += 1
        medio = (izq + der) // 2
        if datos[medio] == buscar:
            print('BUSQUEDA BINARIA: Número encontrado en la posición',
                  medio, 'Iterando', iteraciones, 'veces.')
            break
        if datos[medio] > buscar:
            der = medio - 1
        if datos[medio] < buscar:
            izq = medio + 1
    if buscar != datos[medio]:
        print('BUSQUEDA BINARIA: No se encontró el número dentro del array.')


# ALGORITMO DE BUSQUEDA FIBONACCI
def BusquedaFibonacci(datos, buscar):
    iteraciones = 0
    exit = 0
    fibM_minus_2 = 0
    fibM_minus_1 = 1
    fibM = fibM_minus_1 + fibM_minus_2
    while (fibM < len(datos)):
        iteraciones = 0
        fibM_minus_2 = fibM_minus_1
        fibM_minus_1 = fibM
        fibM = fibM_minus_1 + fibM_minus_2
    index = -1
    while (fibM > 1):
        i = min(index + fibM_minus_2, (len(datos)-1))
        if (datos[i] < buscar):
            iteraciones += 1
            fibM = fibM_minus_1
            fibM_minus_1 = fibM_minus_2
            fibM_minus_2 = fibM - fibM_minus_1
            index = i
        elif (datos[i] > buscar):
            iteraciones += 1
            fibM = fibM_minus_2
            fibM_minus_1 = fibM_minus_1 - fibM_minus_2
            fibM_minus_2 = fibM - fibM_minus_1
        else:
            print('BUSQUEDA FIBONACCI: Número encontrado en la posición',
                  i, 'Iterando', iteraciones, 'veces.')
            exit = 1
            break
    if (fibM_minus_1 and index < (len(datos)-1) and datos[index+1] == buscar and exit != 1):
        print('BUSQUEDA FIBONACCI: Número encontrado en la posición', index + 1)
    elif (exit != 1):
        print('BUSQUEDA FIBONACCI: No se encontró el número dentro del array.')


# =====================================
# INICIO FUNCIONES PARA GENERAR ARRAY
generarArrayDesordenado()
# generarArrayOrdenado_Mm()
# generarArrayOrdenado_mM()
# generarArrayParcial()
# FIN FUNCIONES PARA GENERAR ARRAY
# =====================================

print()
print()

# =========================
# INICIO ARRAY DESORDENADO
print("ARRAY COMPLETAMENTE DESORDENADO:")
print(numeros)
Quicksort(numeros, 0, len(numeros) - 1)
print()
print("ARRAY ORDENADO POR QUICKSORT:")
print(numeros)
print()
print()

search = random.choice(numeros)
print('Valor a buscar en array ordenado por QUICKSORT mediante BUSQUEDA LINEAL:', search)
busquedaLineal(numeros, search)
print()
print('Valor a buscar en array ordenado por QUICKSORT mediante BUSQUEDA BINARIA:', search)
busquedaBinaria(numeros, search)
print()
print('Valor a buscar en array ordenado por QUICKSORT mediante BUSQUEDA FIBONACCI:', search)
BusquedaFibonacci(numeros, search)
# FIN ARRAY DESORDENADO
# =========================

print()
print()
print()
