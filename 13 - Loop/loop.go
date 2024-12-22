package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		fmt.Println("Incrementando i", i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
	}

	nomes := [3]string{"Jonas", "Carlos", "Amanda"}

	for index, name := range nomes {
		fmt.Println(index, name)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"nome":      "Douglas",
		"sobrenome": "Silva",
	}

	for key, value := range usuario {
		fmt.Println(key, value)
	}

	type userStruct struct {
		nome      string
		sobrenome string
	}

	usuarios := [2]userStruct{{"Douglas", "Silva"}, {"Giba", "Mina"}}

	for k := 0; k < len(usuarios); k++ {
		fmt.Println(usuarios[k])
	}

	/* Loop INFINITO
	for {
	}*/
}
