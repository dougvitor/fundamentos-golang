package main

import "fmt"

func main() {

	usuario1 := map[string]string{
		"nome":      "Douglas",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario1)

	usuario2 := map[string]map[string]string{
		"dados_pessoais": usuario1,
		"curso": {
			"nome":        "Ciência da computação",
			"instituicao": "Universidade Paulista - UNIP",
		},
		"signo": {
			"nome": "Áries",
		},
	}

	fmt.Println(usuario2)

	fmt.Println(usuario2["dados_pessoais"]["nome"], usuario2["curso"]["nome"], usuario2["curso"]["instituicao"], usuario2["signo"]["nome"])

	usuario2["torcedor"] = map[string]string{
		"time": "Palmeiras",
	}

	fmt.Println(usuario2)

	delete(usuario1, "sobrenome")

	fmt.Println(usuario2)

	delete(usuario2, "dados_pessoais")

	fmt.Println(usuario1)
	fmt.Println(usuario2)

}
