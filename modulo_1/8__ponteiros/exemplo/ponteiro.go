package main

import "fmt"

func main() {
	x := "Oi"

	var p *string = &x // Ponteiro para a variável x

	fmt.Println("Valor de x:", x)
	fmt.Println("Endereço de x:", &x)

	fmt.Println("Valor de p:", *p) // Desreferenciando o ponteiro para obter o valor
	fmt.Println("Endereço armazenado em p:", p)
	fmt.Println("Endereço de p:", &p) // Endereço do ponteiro
}
