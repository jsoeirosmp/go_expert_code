package main

import (
	"fmt"
)

// Maps tem a estrutura de key e value, sendo ela map[type]type{"key":value}
func main() {
	salarios := map[string]int{"Joao": 1000, "Wesley": 1500, "Maria": 2000}

	// Printando o map inteiro
	fmt.Println(salarios)
	// Printando somente uma key, no caso um nome
	fmt.Printf("O salario de João é: %d\n", salarios["Joao"])

	// Deletando uma chave
	delete(salarios, "Maria")
	fmt.Println("Salario de Maria deletado: ", salarios)

	// Criando uma chave
	salarios["Paulo"] = 1700
	fmt.Println("Salario de Paulo adicionado:", salarios)

	// Posso usar a func make() pra estruturar o map vazio como na forma abaixo, e depois ir manipulando ele
	// meuMap := make(map[string]int)
	// meuMap := map[string]int{}

	// Iterando pelo map baseado no range dele usando a chave e o valor
	for nome, salario := range salarios {
		fmt.Printf("Nome: %v\n Salário: %d\n", nome, salario)
	}

	// Usando blank identifier pra ignorar variavel que nao quero utilizar, por exemplo
	for _, salario := range salarios {
		fmt.Printf("Salário: %d\n", salario)

	}
}
