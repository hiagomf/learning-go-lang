package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

// https://medium.com/gommunity/operadores-c7e1a41cfd4
// Definições de variáveis com VAR, considera-se a mesma regra nos parâmetros de funções, considerando-se o último passado
var c, python, java bool
var j, k int = 2, 3

func main() {
	fmt.Println("Hello world")
	fmt.Println("Meu número favorito é:", rand.Intn(10))
	fmt.Println("O valor de PI é: ", math.Pi)
	fmt.Println("A soma de dois números é: ", add(2, 3))

	//Criando e setando valores
	a, b := swap("Palavra 1", "Palavra 2")
	fmt.Println(a, b)

	fmt.Println(split(17))

	printGlobalVars()
	conversion()
	constants()
	sumWithFor()
	sumWithoutFirstAndLastParams()
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	usingSwitch()
	usingSwitchDays()
	ifThenElse()
	retardExecution()
}

//Caso haja tipo de parâmetros, tipa-se após a variável, apenas a ultima seja tipada, considera-se o tipo dela
// ou seja, caso tipo diferente declara em todos, caso não, apenas no último
func add(x int, y int) int {
	return x + y
}

//Retornar múltiplos resultados
func swap(x, y string) (string, string) {
	return y, x
}

//Em casos de retornos limpos, define-se o retorno no início da função, dentro do segundo parêntesis
func split(sum int) (x, y int) {
	x = sum * 4
	y = sum - x
	return
}

func conversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("Números convertidos: ", x, y, z)
}

func printGlobalVars() {
	fmt.Println("Variáveis globais ", c, python, java, j, k)
}

func constants() {
	const Pi = 3.14

	const World = "Mendes"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func sumWithFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//Não há while em GO
func sumWithoutFirstAndLastParams() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func usingSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func usingSwitchDays() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("Today is ", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func ifThenElse() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func retardExecution() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
