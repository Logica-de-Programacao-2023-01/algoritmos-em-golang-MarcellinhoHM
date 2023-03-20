package main

import "fmt"

// Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.
func main() {
	var x int
	fmt.Println("Digite um número")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Printf("%d é par", x)
	} else {
		fmt.Printf("%d é ímpar")
	}

}
