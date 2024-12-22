package main

import (
	"fmt"
	enderecos "test_module/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Rua Itaúna")
	fmt.Println(tipoEndereco)
}
