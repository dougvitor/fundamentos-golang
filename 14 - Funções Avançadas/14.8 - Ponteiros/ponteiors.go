package main

import "fmt"

func inverterSinalComCopiaDeValor(numero int) int {
	return numero * -1
}

func inverterSinalComReferencia(numero *int) {
	*numero = *numero * -1
}

func main() {

	numero := 20
	numeroInvertido := inverterSinalComCopiaDeValor(numero)
	fmt.Println(numero)
	fmt.Println(numeroInvertido)

	outroNumero := 40
	fmt.Println(outroNumero)
	inverterSinalComReferencia(&outroNumero)
	fmt.Println(outroNumero)
}
