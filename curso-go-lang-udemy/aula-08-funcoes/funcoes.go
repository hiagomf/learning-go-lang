package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	//Declarando função como variável
	var f = func(txt string) string {
		return txt
	}

	retornoFuncaoF := f("Texto da função F")
	fmt.Println(retornoFuncaoF)

	// Caso não vá usar, substitui o nome da variável pelo underline, mas é necessário as duas instâncias
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
