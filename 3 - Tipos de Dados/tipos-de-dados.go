package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 1000000000
	fmt.Println(numero3)

	// uint8 = byte
	var numero4 byte = 100
	fmt.Println(numero4)

	var numeroReal1 float32 = 100.01
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1000000000000000000000000000000000000000000000000.01
	fmt.Println(numeroReal2)

	numeroReal3 := 1000000.01
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	char := 'D'
	fmt.Println(char)

	var numero5 int
	fmt.Println(numero5)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("Internal Error")
	fmt.Println(erro2)

}
