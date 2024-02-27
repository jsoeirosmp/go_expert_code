package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	Adress Endereco
}

// Método pra alterar variavel dentro da struct
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O user %s foi desativado, Status atual: %v", c.Nome, c.Ativo)
}
func (c Cliente) Ativar() {
	c.Ativo = true
	fmt.Printf("O user %s foi ativado, Status atual: %v", c.Nome, c.Ativo)
}

func main() {
	joao := Cliente{
		Nome:  "Joao",
		Idade: 26,
		Ativo: true,
	}
	joao.Desativar()
	joao.Ativar()
}
