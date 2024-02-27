package main

import "fmt"

func main() {
	// Declarando array de 3 posiçoes, tamanho fixo
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	fmt.Println("Meu array contem os elementos:", meuArray)

	// Printando ultima posicao do array, se a len é X e comeca de 0, sempre vai ser len-1
	fmt.Println("A ultima posiçao do meu array é:", len(meuArray)-1)

	// Printando ultimo elemento
	fmt.Println("O último elemento do meu array é:", meuArray[len(meuArray)-1])

	// Printando elemento em especifico
	fmt.Println("Printando elemento na posiçao 0:", meuArray[0])

	// Iterando o array, nao preciso inicializar as vars pois o Go ja infere elas como 0, entao ja começo a iterar
	for indice, valor := range meuArray {
		fmt.Printf("O valor da posiçao %d é %d\n", indice, valor)
	}

}
