package carquitetura

import (
	"runtime"
	"testing"
)

// go test -v
// cd 8-Teste -> go test ./...

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}

	t.Fail()
}
