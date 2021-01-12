package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero < 15 {
		fmt.Println("Menor que 15")
	} else {
		fmt.Println("Igual a 15")
	}

	//if init é a criação da variável no if, mas só poder ser utilizada dentro do escopo
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if outroNumero == 0 {
		fmt.Println("Número igual a zero")
	} else {
		fmt.Println("Número é menor que zero")
	}
}
