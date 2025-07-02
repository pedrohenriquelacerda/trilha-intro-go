package main

// Crie uma variável 'nota', logo após crie uma condição para aprovar alunos
// 'nota' maior ou igual a 7 imprime "Aprovado", caso contrário imprime "Reprovado"

func main() {
	var nota float64 = 10

	if nota >= 7 {
		println("Aprovado")
	} else {
		println("Reprovado")
	}

	switch nota {
	case 10:
		println("Aprovadasso")
	case 9, 8:
		println("Aprovado com nota alta")
	case 7:
		println("Aprovado")
	default:
		println("Reprovado")
	}
}
