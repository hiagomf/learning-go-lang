package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Map é basicamente um array que pode ser personalizável tendo seus tipos: [key]valor
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "da Silva",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "CRAJUBAR",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "curso")
	fmt.Println("APÓS APAGADO ", usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Touro",
	}
	fmt.Println("COM SIGNO", usuario2)
}
