package main

import "fmt"

func main() {
	fmt.Println(soma(1, 5))
}

// Quero somar uma qtd de numeros que nao sei qual é exatamente, pra isso uso as variadics, nesse caso os ...
// que indicam que vao ser varios valores
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
