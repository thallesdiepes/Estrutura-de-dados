package main

import "fmt"

func main() {
	var x int
	var y float64
	var msg string

	fmt.Println("Informe um número:")
	fmt.Scan(&x)

	fmt.Println("Informe outro número:")
	fmt.Scan(&y)

	fmt.Println("Informe um texto:")
	fmt.Scan(&msg)

	fmt.Println(x, y, msg)
}
