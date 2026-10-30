package main

import "fmt"

func swap(a *int, b *int) {
	aux := *a
	*a = *b
	*b = aux
}

func sumArray(values []int) int {
	suma := 0
	for i := 0; i < len(values); i++ {
		suma += values[i]
	}
	return suma
}

func minArray(values []int) int {
	menor := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < menor {
			menor = values[i]
		}
	}
	return menor
}

func maxArray(values []int) int {
	mayor := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > mayor {
			mayor = values[i]
		}
	}
	return mayor
}

func countNegatives(values []int) int {
	cantidad := 0
	for i := 0; i < len(values); i++ {
		if values[i] < 0 {
			cantidad++
		}
	}
	return cantidad
}

func printReverse(values []int) {
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Print(values[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func copyArray(values []int) []int {
	// Se crea otro slice para no cambiar el original.
	copia := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		copia[i] = values[i]
	}
	return copia
}

func sortAscending(values []int) {
	// Ordenamiento de burbuja.
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				aux := values[j]
				values[j] = values[j+1]
				values[j+1] = aux
			}
		}
	}
}

func findPair(values []int, target int) bool {
	izquierda := 0
	derecha := len(values) - 1

	for izquierda < derecha {
		suma := values[izquierda] + values[derecha]
		if suma == target {
			fmt.Printf("Par encontrado: (%d, %d)\n", values[izquierda], values[derecha])
			return true
		} else if suma < target {
			izquierda++
		} else {
			derecha--
		}
	}
	return false
}

func threeSum(values []int) int {
	if len(values) < 3 {
		return 0
	}

	copia := copyArray(values)
	sortAscending(copia)
	cantidad := 0

	for i := 0; i < len(copia)-2; i++ {
		if i > 0 && copia[i] == copia[i-1] {
			continue
		}

		izquierda := i + 1
		derecha := len(copia) - 1
		for izquierda < derecha {
			suma := copia[i] + copia[izquierda] + copia[derecha]

			if suma == 0 {
				fmt.Printf("Terna: (%d, %d, %d)\n", copia[i], copia[izquierda], copia[derecha])
				cantidad++
				izquierda++
				derecha--

				// Se saltan los números repetidos.
				for izquierda < derecha && copia[izquierda] == copia[izquierda-1] {
					izquierda++
				}
				for izquierda < derecha && copia[derecha] == copia[derecha+1] {
					derecha--
				}
			} else if suma < 0 {
				izquierda++
			} else {
				derecha--
			}
		}
	}
	return cantidad
}

func buildLeftMax(height []int, leftMax []int) {
	if len(height) == 0 {
		return
	}

	mayor := height[0]
	for i := 0; i < len(height); i++ {
		if height[i] > mayor {
			mayor = height[i]
		}
		leftMax[i] = mayor
	}
}

func buildRightMax(height []int, rightMax []int) {
	if len(height) == 0 {
		return
	}

	mayor := height[len(height)-1]
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > mayor {
			mayor = height[i]
		}
		rightMax[i] = mayor
	}
}

func trappedWater(height []int) int64 {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)

	total := int64(0)
	for i := 0; i < len(height); i++ {
		limite := leftMax[i]
		if rightMax[i] < limite {
			limite = rightMax[i]
		}
		total += int64(limite - height[i])
	}
	return total
}

func mostrarDetalleAgua(height []int) {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	agua := make([]int, len(height))
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)

	for i := 0; i < len(height); i++ {
		limite := leftMax[i]
		if rightMax[i] < limite {
			limite = rightMax[i]
		}
		agua[i] = limite - height[i]
	}

	fmt.Println("Alturas:      ", height)
	fmt.Println("Max izquierda:", leftMax)
	fmt.Println("Max derecha:  ", rightMax)
	fmt.Println("Agua:         ", agua)
}

func probarThreeSum(values []int, esperado int) {
	fmt.Println("Entrada:", values)
	cantidad := threeSum(values)
	fmt.Println("Cantidad:", cantidad, "| Esperado:", esperado)
	fmt.Println()
}

func probarAgua(height []int, esperado int64) {
	resultado := trappedWater(height)
	fmt.Println("Alturas:", height, "| Agua:", resultado, "| Esperado:", esperado)
}

func main() {
	fmt.Println("EJERCICIO 1 - Intercambio")
	a := 8
	b := -3
	fmt.Println("Antes:", a, b)
	swap(&a, &b)
	fmt.Println("Despues:", a, b)
	fmt.Println()

	values := []int{4, -2, 0, 7, -5}
	fmt.Println("EJERCICIO 2 - Estadisticas")
	fmt.Println("Arreglo:", values)
	fmt.Println("Suma:", sumArray(values))
	fmt.Println("Minimo:", minArray(values))
	fmt.Println("Maximo:", maxArray(values))
	fmt.Println("Negativos:", countNegatives(values))
	fmt.Print("Al reves: ")
	printReverse(values)
	fmt.Println()

	fmt.Println("EJERCICIO 3 - Copia")
	copia := copyArray(values)
	copia[0] = 99
	fmt.Println("Original:", values)
	fmt.Println("Copia:", copia)
	fmt.Println()

	fmt.Println("EJERCICIO 4 - Ordenamiento")
	ordenado := copyArray(values)
	sortAscending(ordenado)
	fmt.Println("Original:", values)
	fmt.Println("Ordenado:", ordenado)
	fmt.Println()

	fmt.Println("EJERCICIO 5 - Buscar pares")
	pares := []int{-4, -1, 0, 2, 5}
	fmt.Println("Entrada:", pares, "Objetivo: 1")
	fmt.Println("Existe:", findPair(pares, 1))
	fmt.Println("Entrada:", pares, "Objetivo: 8")
	fmt.Println("Existe:", findPair(pares, 8))
	fmt.Println("Entrada: [0 0] Objetivo: 0")
	fmt.Println("Existe:", findPair([]int{0, 0}, 0))
	fmt.Println("Entrada: [3] Objetivo: 6")
	fmt.Println("Existe:", findPair([]int{3}, 6))
	fmt.Println()

	fmt.Println("EJERCICIO 6 - 3Sum")
	probarThreeSum([]int{-1, 0, 1, 2, -1, -4}, 2)
	probarThreeSum([]int{0, 1, 1}, 0)
	probarThreeSum([]int{0, 0, 0}, 1)
	probarThreeSum([]int{0, 0, 0, 0}, 1)
	probarThreeSum([]int{-2, 0, 0, 2, 2}, 1)
	probarThreeSum([]int{0, 0}, 0)

	fmt.Println("EJERCICIO 7 - Maximos")
	height := []int{3, 0, 2, 0, 4}
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)
	fmt.Println("Alturas:", height)
	fmt.Println("Izquierda:", leftMax)
	fmt.Println("Derecha:", rightMax)
	fmt.Println()

	fmt.Println("EJERCICIO 8 - Agua atrapada")
	probarAgua([]int{5}, 0)
	probarAgua([]int{0, 1, 2, 3}, 0)
	probarAgua([]int{3, 2, 1, 0}, 0)
	probarAgua([]int{2, 2, 2}, 0)
	probarAgua([]int{3, 0, 3}, 3)
	probarAgua([]int{3, 0, 2, 0, 4}, 7)
	probarAgua([]int{4, 2, 0, 3, 2, 5}, 9)
	probarAgua([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)
	fmt.Println("Caso propio:")
	probarAgua([]int{4, 1, 3, 0, 4}, 8)
	fmt.Println("Detalle de una prueba:")
	mostrarDetalleAgua([]int{3, 0, 2, 0, 4})
}
