package main

import "fmt"

func main() {

	numero := 15

	if numero > 10 {
		fmt.Println("Número é maior que dez")
	} else {
		fmt.Println("Número é menor que dez")
	}

	numero2 := -7

	if outroNumero := numero2; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if outroNumero < 10 {
		fmt.Println("Numero é menor que dez")
	} else {
		fmt.Println("Número está entre zero e dez")
	}

}
