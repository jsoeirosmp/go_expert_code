package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// A interface é um conjunto de métodos com o qual as structs podem se conformar, neste caso, todo cliente
// pode ser ativada ou desativada, entao adiciono ele a essa interface Pessoa, que contem metodos comuns a
// todos que eu considero "pessoas". Lembrar exemplo interface animal Sound(), struct dog, dog entra na interface
// animal pois todo animal emite som, interfaces aceitam SOMENTE METODOS, e nao outras propriedades como "nome:"
type Pessoa interface {
	Desativar()
}

// Criando segunda struct
type Empresa struct {
	Nome string
}

// Func somente com metodo desativar
func (e Empresa) Desativar() {

}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	Adress Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O user %s foi desativado, Status atual: %v", c.Nome, c.Ativo)
}

// Criando func pra poder mexer com as variaveis que possuem a interface pessoa
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()

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
	// Com a interface e a func criadas, posso desativar o user de maneira mais simples
	// Pq posso passar joao se o type recebido dessa func é uma pessoa? Pq a var joao é considerada uma pessoa
	// por implementar esse metodo desativar
	Desativacao(joao)

	// Declaro a var de empresa vazia so pra teste, e ao usar o metodo Desativacao(), funciona do mesmo jeito
	// porque fiz aquela func recebendo empresa que tem o metodo Desativar, ou seja, TODA struct com o metodo
	// desativar sera considerada uma "pessoa" pois vai implementar a interface
	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)
}
