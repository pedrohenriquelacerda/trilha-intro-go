package main

import (
	"fmt"
	"time"
)

// Exemplo dos slides

func sayHello() {
	fmt.Println("Olá de uma goroutine!")
}

func main() {
	go sayHello() // inicia uma goroutine
	fmt.Println("Olá do main!")
	time.Sleep(2 * time.Second) // espera para a goroutine executar
	fmt.Println("Encerrando a execução")
}
