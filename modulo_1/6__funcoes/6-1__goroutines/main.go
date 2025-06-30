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
	go countTo10()
	sayHello()
	time.Sleep(100 * time.Second)
}

func countTo10() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
