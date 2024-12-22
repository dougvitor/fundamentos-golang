package main

import "fmt"

func isAlunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	} else {
		panic("A MÉDIA É EXATAMENTE 6!")
	}
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func main() {

	fmt.Println(isAlunoAprovado(3, 6))
	fmt.Println("Pós execução!")

}
