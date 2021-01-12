package main

import "fmt"

//Structures são coleções de campos, são quase como um campo que armazena dados variados de tipos diferentes, como na orientação a objeto
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo struct")

	var u usuario
	u.nome = "Hiago"
	u.idade = 21
	fmt.Println(u)

	u2 := usuario{"Mendes", 23, endereco{"Rua sta Isabel", 12}}
	fmt.Println(u2)

	u3 := usuario{idade: 27}
	fmt.Println(u3)
}
