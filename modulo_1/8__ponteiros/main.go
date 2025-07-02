package main

// 1. Crie uma struct chamada Usuario com os campos Nome (string) e Idade (int).
// 2. Crie uma função chamada AtualizarUsuario que recebe um ponteiro para Usuario e dois parâmetros: um novo nome e uma nova idade.
// 3. A função deve alterar os campos da struct usando o ponteiro.
// 4. No main(), crie um usuário com nome "João" e idade 25.
// 5. Imprima o valor do usuário antes e depois de chamar AtualizarUsuario, passando o novo nome "Carlos" e idade 30.

type Usuario struct {
	Nome  string
	Idade int
}

func AtualizarUsuario(u *Usuario, novoNome string, novaIdade int) {
	u.Nome = novoNome
	u.Idade = novaIdade
}

func main() {
	usuario := Usuario{
		Nome:  "João",
		Idade: 25,
	}

	println("Antes da atualização:")
	println("Nome:", usuario.Nome)
	println("Idade:", usuario.Idade)

	AtualizarUsuario(&usuario, "Carlos", 30)

	println("Depois da atualização:")
	println("Nome:", usuario.Nome)
	println("Idade:", usuario.Idade)
}
