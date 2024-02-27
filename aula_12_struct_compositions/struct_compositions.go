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
	//Chamando a struct de cima aqui dentro, fazendo uma "composicao"
	Endereco
	// Declarando um campo do type Endereco
	Adress Endereco
}

func main() {
	// Criando var com dados baseados na struct
	joao := Cliente{
		Nome:  "Joao",
		Idade: 26,
		Ativo: true,
	}
	// Alterando valores da struct endereco dentro da Cliente, posso fazer com joao.Endereco.Cidade tbm
	joao.Cidade = "Sao Jose dos Campos"
	fmt.Println(joao)

	// Alterando o Adress que tem o type que eu criei (notar que agora ele esta com o valor acima e tambem o SJC
	// na struct, pois joao.Cidade =/= joao.Adress.Cidade)
	joao.Adress.Cidade = "SJC"
	fmt.Println(joao)
}
