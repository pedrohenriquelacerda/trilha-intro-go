package main

// Inicialize um array com 3 elementos, utilize o for() clássico para
// imprimir todos os elementos na tela e após isso faça a mesma coisa com for-range()

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	// Usando for clássico
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	// Usando for-range
	for index, value := range arr {
		fmt.Printf("O indíce é %d e o valor é %d\n", index, value)
	}
}
