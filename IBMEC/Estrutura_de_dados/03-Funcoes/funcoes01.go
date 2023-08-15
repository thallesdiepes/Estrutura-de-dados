package main

import "fmt"

func main() {
	fmt.Println(soma(2, 6))
	fmt.Println(subtracao(4, 7))

	usoAnonima()
}

//funções anonimas
func usoAnonima() {
	dobra := func(x int) int {
		return x * 2
	}

	resultado := dobra(5)
	fmy.Println(resultado)
}

//func nome(parametro1 tipo,parametro2 tipo, ...) tipoRetorno {}
func soma(a int, b int) int {
	return a + b
}

//Parametros com tipo repetido podem ser omitidos
func subtracao(a, b float64) float64 {
	return a - b
}
