package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var variavel3 int = 100
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	fmt.Println(variavel3, *ponteiro)

	var variavel4 int = *ponteiro
	fmt.Println(variavel3, variavel4)

	variavel3 = 150
	fmt.Println(variavel3, variavel4)

	fmt.Println(variavel3, *ponteiro)

	ponteiro2 := &ponteiro

	fmt.Println(ponteiro, ponteiro2)

	fmt.Println(*ponteiro, *ponteiro2)

}
