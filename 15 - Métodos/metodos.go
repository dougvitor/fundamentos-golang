package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) isMaiorIdade() bool {
	return u.idade >= 18
}

func (u *usuario) incrementarIdade() {
	u.idade++
}

func main() {

	user1 := usuario{"Douglas", 41}
	user1.salvar()

	user2 := user1
	fmt.Println(user2.isMaiorIdade())
	fmt.Println(user2.idade)

	user2.incrementarIdade()
	fmt.Println(user2.idade)

	fmt.Println(user1.idade)

}
