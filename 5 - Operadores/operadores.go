package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println("-----------------------------")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	//NÃO EXISTE OPERADOR TERNÁRIO

}
