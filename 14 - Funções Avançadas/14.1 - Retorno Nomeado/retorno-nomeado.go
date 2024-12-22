package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int, multiplicacao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	multiplicacao = n1 * n2
	return
}

func main() {
	soma, subtracao, multiplicacao1 := calculosMatematicos(25, 80)
	fmt.Println(soma, subtracao, multiplicacao1)
}
