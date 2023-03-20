package main

import "fmt"

// Faça um algoritmo que leia três números inteiros e mostre o menor deles.
func main() {
	var x int
	var y int
	var z int
	fmt.Print("Qual o primeiro número?")
	fmt.Scan(&x)
	fmt.Println("Qual o segundo número?")
	fmt.Scan(&y)
	fmt.Println("Qual o terceiro número?")
	fmt.Scan(&z)
	if x < y && x < z {
		fmt.Println("O primeiro é o menor")
	} else if y < x && y < z {
		fmt.Println("O segundo é o menor")
	} else if z < x && z < y {
		fmt.Println("O terceiro é o menor")
	} else {
		fmt.Println("Todos são iguais")
	}

}
