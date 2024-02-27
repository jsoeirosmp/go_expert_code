package main

import "fmt"

// Closure é uma func "anonima" dentro de outra func, nesse caso estou pegando um valor pronto e trabalhando
// ele pra exibir outro resultado
func main() {
	total := func() int {
		return soma(1, 5) * 2
	}()
	fmt.Println(total)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
