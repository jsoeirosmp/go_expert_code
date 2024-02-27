package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	// Criando var com dados baseados na struct
	joao := Cliente{
		Nome:  "Joao",
		Idade: 26,
		Ativo: true,
	}
	fmt.Println(joao)

	// Alterando valores
	joao.Nome = "Joao Vitor"
	joao.Idade = 30
	joao.Ativo = false
	fmt.Println(joao)
}
