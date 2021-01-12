package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo o arquivo main")
	auxiliar.Escrever()

	//Utilizando função de pacote adicionado com a linha de comando go get github.com/badoux/checkmail
	//O comando go tidy remove todas as dependências que não estão sendo utilizadas.
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
