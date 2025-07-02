package main

// Crie um slice de strings com 3 nomes e adicione um quarto nome usando append()
// e imprima todos os nomes.

// EXTRA: Imprimir na tela o cap() e len() do slice antes de adicionar o quarto nome
// e depois de adiciona-lo, observar a diferença

import "fmt"

func main() {
	var nomes = []string{"Alice", "Bob", "Charlie"}

	fmt.Printf("Antes de adicionar: cap=%d, len=%d\n", cap(nomes), len(nomes))

	nomes = append(nomes, "David")

	fmt.Printf("Depois de adicionar: cap=%d, len=%d\n", cap(nomes), len(nomes))

	for _, nome := range nomes {
		fmt.Println(nome)
	}
}
