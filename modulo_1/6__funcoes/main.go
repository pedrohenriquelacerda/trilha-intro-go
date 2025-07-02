package main

// 1. Crie uma função que faça a divisão de dois números e diga se o resultado é par ou ímpar
// 2. A função deve ter 2 retornos, o resultado da divisão e a indicação se o número é par ou ímpar (string)
// 3. Chame a função dentro da main() e imprima o resultado do retorno

func divisaoParOuImpar(num1, num2 int) (int, string) {
	if num2 == 0 {
		return 0, "Divisão por zero não é permitida"
	}

	resultado := int(num1) / int(num2)

	if int(resultado)%2 == 0 {
		return resultado, "par"
	}

	return resultado, "ímpar"
}

func main() {
	num1 := 10
	num2 := 3

	resultado, parOuImpar := divisaoParOuImpar(num1, num2)
	println("Resultado da divisão:", resultado)
	println("O resultado é:", parOuImpar)
}
