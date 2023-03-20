package main

import "fmt"

// Faça um algoritmo que imprima os números de 1 a 10 em ordem crescente.
func main() {

	for x := 0; x < 15; x++ {
		if x == 11 {
			break

		}
		fmt.Println(x)

	}
	fmt.Println("Fim do loop com break")

}
