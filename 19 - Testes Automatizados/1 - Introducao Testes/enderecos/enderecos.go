package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoLetraMinuscula, " ")[0]

	c := cases.Title(language.Und)

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			return c.String(primeiraPalavraEndereco)
		}
	}

	return "Tipo inválido!"
}
