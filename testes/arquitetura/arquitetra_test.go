package arquitetura

import "testing"
import "runtime"

func TestDependente(t *testing.T) {
	t.Parallel() // serve para fazer testes em paralelo.
	// Ajuda em casos onde há muitos testes.
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona na arquitetura amd64")
	}
	//...
	t.Log("Falhou miseravelmente")
	t.Fail()
}
