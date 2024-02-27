package main

import "fmt"

func main() {
	// Declarando slice
	meuSlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 150}
	fmt.Printf("Len = %d\n Cap = %d\n Items = %v", len(meuSlice), cap(meuSlice), meuSlice)

	// Ignorando tudo que estiver a direita da posicao zero, slice ficará vazio
	fmt.Printf("Len = %d\n Cap = %d\n Items = %v", len(meuSlice[:0]), cap(meuSlice[:0]), meuSlice[:0])

	// Ignorando tudo que estiver a direita da P4, nesse caso, mostra só os elementos ate o "40"
	fmt.Printf("Len = %d\n Cap = %d\n Items = %v", len(meuSlice[:4]), cap(meuSlice[:4]), meuSlice[:4])

	// Ignorando tudo que estiver a esquerda da P6, printando do 70 ate o final
	fmt.Printf("Len = %d\n Cap = %d\n Items = %v", len(meuSlice[6:]), cap(meuSlice[6:]), meuSlice[6:])

	// Aumentando slice, vai dobrar de tamanho quando bater a max capacity
	meuSlice = append(meuSlice, 160)
	fmt.Printf("Len = %d\n Cap = %d\n Items = %v", len(meuSlice), cap(meuSlice), meuSlice)

}
