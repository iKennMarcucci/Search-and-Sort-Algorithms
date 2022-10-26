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


# ALGORITMOS DE ORDENAMIENTO POR BURBUJA
def Burbuja(array):
    switch = True
    while switch:
        switch = False
        for i in range(len(array) - 1):
            if array[i] > array[i + 1]:
                array[i], array[i + 1] = array[i + 1], array[i]
                switch = True


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
Burbuja(numeros)
print()
print("ARRAY ORDENADO POR BURBUJA:")
print(numeros)
print()
print()

search = random.choice(numeros)
print('Valor a buscar en array ordenado por BURBUJA mediante BUSQUEDA LINEAL:', search)
busquedaLineal(numeros, search)
print()
print('Valor a buscar en array ordenado por BURBUJA mediante BUSQUEDA BINARIA:', search)
busquedaBinaria(numeros, search)
print()
print('Valor a buscar en array ordenado por BURBUJA mediante BUSQUEDA FIBONACCI:', search)
BusquedaFibonacci(numeros, search)
# FIN ARRAY DESORDENADO
# =========================

print()
print()
print()
