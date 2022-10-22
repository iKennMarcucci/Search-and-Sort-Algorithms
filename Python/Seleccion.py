import random

numeros = []

# GENERADOR DE ARRAYS COMPRENDIDO POR EL TAMAÑO


def generarArrayLongitud100():
    # 10^2
    for i in range(100):
        num = random.randint(10000, 99999)
        numeros.append(num)


generarArrayLongitud100()

# ALGORITMOS DE ORDENAMIENTO POR SELECCION


def seleccion(array):
    longitud = len(array)
    for i in range(longitud - 1):
        for j in range(i + 1, longitud):
            if array[i] > array[j]:
                array[i], array[j] = array[j], array[i]


print("Array Original: ")
print(numeros)
seleccion(numeros)
print("Array Ordenado: ")
print(numeros)
print()


# ALGORITMO DE BUSQUEDA LINEAL
def busquedaLineal(datos, buscar):
    i = 0
    for x in datos:
        if x == buscar:
            print('BUSQUEDA LINEAL. Número encontrado en la posición: ', i)
            break
        i += 1
    if i == len(datos):
        print('BUSQUEDA LINEAL. No se encontró el número dentro del array.')


search = random.choice(numeros)
print('Valor a buscar: ', search)
busquedaLineal(numeros, search)
print()
