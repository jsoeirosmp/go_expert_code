package main

import "fmt"

// Globais
const a = "Hello, world!"

// Criando um type, no caso ID
type ID int

var (
	b bool    = true
	c int     = 66
	d string  = "Joao"
	e float64 = 1.8
	f ID      = 1
)

func main() {
	// Local
	var a string = "olar"
	// Declarando e deixando o Go inferir o type
	b := false
	fmt.Println(a, b, f)
	fmt.Printf("O type da var b é: %T", b)
	fmt.Printf("O valor da var c é: %v", c)
}
