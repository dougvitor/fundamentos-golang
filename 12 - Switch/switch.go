package main

import "fmt"

func textoDiaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Dia inválido"
	}
}

func numeroDiaDaSemana(texto string) int {
	var dia int

	switch {
	case texto == "Domingo":
		dia = 1
	case texto == "Segunda":
		dia = 2
	case texto == "Terça":
		dia = 3
	case texto == "Quarta":
		dia = 4
	case texto == "Quinta":
		dia = 5
	case texto == "Sexta":
		dia = 6
	case texto == "Sabado":
		dia = 7
		fallthrough
	default:
		dia = 0
	}

	return dia
}

func main() {

	fmt.Println(textoDiaDaSemana(7))
	fmt.Println(numeroDiaDaSemana(textoDiaDaSemana(7)))
	fmt.Println(textoDiaDaSemana(2))
	fmt.Println(numeroDiaDaSemana(textoDiaDaSemana(2)))

}
