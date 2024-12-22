package main

import "fmt"

var numero int

func init() {
	fmt.Println("Executando a função init")
	numero = 25
}

func main() {
	fmt.Println("Função main sendo executada")

	fmt.Println(numero)
}
