package main

import (
	"errors"
	"fmt"
)

func main() {
	// Testando chamada da func soma
	fmt.Println(soma(1, 50))

	// Declarando vars a serem recebidas na func, comum utilizar o ultimo item pra error em Go, se tiver algo
	// diferente de nil, retorna o erro, se erro for nil, retorna o valor que eu quero.
	valor, err := multiplica(2, 20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}
}

// Func bem padrao como ja conheco, recebe valores, especifica type deles e do return, se nao tiver type do return,
// é pq a func nao retorna nada, sendo mesmo type, poderia deixar "(a, b int)""
func soma(a int, b int) (int, bool) {
	if a+b >= 50 {
		// Retornando o int e o bool que especifiquei anteriormente
		return a + b, true
	} else {
		return a + b, false
	}
}

// Usando errors pra informar alguma excecao ou problema, erro retorna vazio se ter td certo
func multiplica(a, b int) (int, error) {
	if a*b >= 50 {
		return 0, errors.New("Resultado da multiplicaçao maior que 50")
	} else {
		return a * b, nil
	}
}
