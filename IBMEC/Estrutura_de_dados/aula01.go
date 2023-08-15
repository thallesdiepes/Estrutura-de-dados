package main

import "fmt"

func main (){
	//Declaração explícita
	var x, y int
	var texto string

	x, y = 4, 5
	texto = "Olá, mundo"

	fmt.Println(x, y, texto)

	//Declaração implícita
	var z = 4.66

	fmt.Println(z)

	//Declaração curta
	msg := "Olá"
	fmt.Println(msg)
}