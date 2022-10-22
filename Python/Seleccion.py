import random

numeros = []


# GENERADOR DE ARRAYS COMPRENDIDO POR EL TAMAÑO
def generarArrayLongitud100():
    # 10^2
    for i in range(100):
        num = random.randint(10000, 99999)
        numeros.append(num)


# ALGORITMOS DE ORDENAMIENTO POR SELECCION
def seleccion(array):
    longitud = len(array)
    for i in range(longitud - 1):
        for j in range(i + 1, longitud):
            if array[i] > array[j]:
                array[i], array[j] = array[j], array[i]


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
            print('BUSQUEDA FIBONACCI: Número encontrado en la posición', i, 'Iterando', iteraciones, 'veces.')
            exit = 1
            break
    if (fibM_minus_1 and index < (len(datos)-1) and datos[index+1] == buscar and exit != 1):
        print('BUSQUEDA FIBONACCI: Número encontrado en la posición', index + 1)
    elif (exit != 1):
        print('BUSQUEDA FIBONACCI: No se encontró el número dentro del array.')


# ===============================================
# INICIO GENERAR ARRAY COMPLETAMENTE DESORDENADO
generarArrayLongitud100()
# FIN GENERAR ARRAY COMPLETAMENTE DESORDENADO
# ===============================================

print()
print()

# =========================
# INICIO ARRAY DESORDENADO
print("ARRAY COMPLETAMENTE DESORDENADO:")
print(numeros)
seleccion(numeros)
print()
print("ARRAY ORDENADO POR SELECCION:")
print(numeros)
print()
print()

search = random.choice(numeros)
print('Valor a buscar en array ordenado por SELECCION mediante BUSQUEDA LINEAL:', search)
busquedaLineal(numeros, search)
print()
print('Valor a buscar en array ordenado por SELECCION mediante BUSQUEDA BINARIA:', search)
busquedaBinaria(numeros, search)
print()
print('Valor a buscar en array ordenado por SELECCION mediante BUSQUEDA FIBONACCI:', search)
BusquedaFibonacci(numeros, search)
# FIN ARRAY DESORDENADO
# =========================

print()
print()
print()
