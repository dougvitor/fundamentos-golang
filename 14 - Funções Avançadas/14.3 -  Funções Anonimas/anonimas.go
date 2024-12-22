package main

import "fmt"

func main() {
	func() {
		fmt.Println("Função Anonima sem parametros")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Função anonima passando Parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebendo parâmetro -> %s e retornando", texto)
	}("Parâmetro enviado")

	fmt.Println(retorno)

}
