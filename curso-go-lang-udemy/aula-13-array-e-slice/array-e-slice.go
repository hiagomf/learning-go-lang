package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array e slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Pos 1", "Pos 2", "Pos 3", "Pos 4", "Pos 5"}
	fmt.Println(array2)

	slice := []int{10, 2, 3, 49, 5, 84, 6486, 46}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(slice))

	// Adiciona na última posição
	slice = append(slice, 12)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//Arrays internos
	fmt.Println("------------------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho
	fmt.Println(cap(slice3)) // Capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Tamanho
	fmt.Println(cap(slice4)) // Capacidade
}
