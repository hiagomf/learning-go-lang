package main

import "fmt"

func main() {
	var texto1 string = "Variável 1"
	texto2 := "Variável 2"

	fmt.Println(texto1, texto2)

	var (
		texto3 string = "lalalalla"
		texto4 string = "lalalalla"
	)

	fmt.Println(texto3, texto4)

	texto5, texto6 := "lalalalalal", "alallalalalaal"
	fmt.Println(texto5, texto6)
}
