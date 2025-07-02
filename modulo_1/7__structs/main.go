package main

// 1. Crie um struct chamado Pessoa com campos nome (string) e idade (int).
// 2. Crie uma variável do tipo Pessoa, atribua valores e imprima na tela.

type Pessoa struct {
	nome  string
	idade int
}

func main() {

	pessoa := Pessoa{
		nome:  "João",
		idade: 30,
	}

	println("Nome:", pessoa.nome)
	println("Idade:", pessoa.idade)
}
