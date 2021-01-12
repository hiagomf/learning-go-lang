package main

import "fmt"

// LEMBRAR QUE NÃO EXISTE HERANÇA EM GO LANG

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    //Essa linha garante que os campos são "copiados" para essa struct
	matricula string
	curso     string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Alvelino", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "2017408371", "Engenharia"}
	fmt.Println(e1.altura)
}
