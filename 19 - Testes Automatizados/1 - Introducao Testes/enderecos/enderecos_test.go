package enderecos_test

import (
	. "test_module/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	input    string
	expected string
}

func TestTipoEndereco(t *testing.T) {
	/*input := "Avenida Paulista"
	expected := "Rua"
	result := TipoEndereco(input)

	if result != expected {
		t.Errorf("O tipo recebido é diferente do esperado! Expected: %s e Result: %s", expected, result)
	}*/

	t.Parallel()

	cenarios := []cenarioDeTeste{
		{"Rua Itaúna", "Rua"},
		{"RUA ITAUNA", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Praça das Bandeiras", "Tipo inválido!"},
		{" ", "Tipo inválido!"},
	}

	for _, cenario := range cenarios {
		result := TipoEndereco(cenario.input)

		if result != cenario.expected {
			t.Errorf("O tipo recebido é diferente do esperado! Expected: %s e Result: %s", cenario.expected, result)
		}
	}
}

func TestInutil(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste falhou!")
	}
}
