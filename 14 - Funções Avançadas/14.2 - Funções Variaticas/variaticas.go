package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numeros := range numeros {
		fmt.Println(texto, numeros)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 45, 67, 89)
	fmt.Println(totalSoma)

	escrever("Ola Mundo", 19, 27, 33, 56, 98)
}
