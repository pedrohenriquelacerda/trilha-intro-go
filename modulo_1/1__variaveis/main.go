package main

// Crie uma variável chamada nome com seu nome e outra idade com sua idade.
// Imprima a frase: "Meu nome é [nome] e tenho [idade] anos."

import "fmt"

func main() {
	var nome string = "Pedro Silva"
	var idade int = 19

	fmt.Printf("Meu nome é %s e tenho %d anos\n", nome, idade)
}
