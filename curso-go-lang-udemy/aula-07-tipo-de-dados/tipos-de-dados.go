package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos de inteiros int8, int16, int32, int64, por padrão caso vazio considera-se 0
	var numero1 int16 = 200
	fmt.Println(numero1)

	// Tipo para números sem sinal uint8, uint16, uint32, uint64, por padrão caso vazio considera-se 0
	var numero2 uint8 = 2
	fmt.Println(numero2)

	//Tipo para números float float32, float64, por padrão caso vazio considera-se 0
	var numero3 float32 = 23442.57
	fmt.Println(numero3)

	//Tipo para string, por padrão caso vazio considera-se vazio
	var str string = "Texto"
	fmt.Println(str)

	// o char:=q retorna o número do caractere da tabela ASCII
	char := 'a'
	fmt.Println(char)

	//Tipo para boleanos
	var booleano1 bool
	fmt.Println(booleano1)

	//Tipo para erros
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
