package strings_test

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperando (%d) <> encontrado (%d). "

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Abner é legal", "Abner", 0},
		{"", "", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
