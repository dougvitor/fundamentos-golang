package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Douglas"
	u.idade = 42
	fmt.Println(u)

	u2 := usuario{"Douglas", 42, endereco{"Rua das Antas", 50}}
	fmt.Println(u2)

	u3 := usuario{idade: 42}
	fmt.Println(u3)
}
