package main

import "fmt"

func main () {
	//array é uma coleção de dados do mesmo tipo, com tamanho definido em compilação

	var filmes [5]string

	filmes[0] = "Shrek"
	filmes[1] = "Star wars"
	fmt.Println(filmes)

	//Declaração curta
	numeros := [4]int{0,2,4,6}

	fmt.Println(numeros)

	//Slices são segmentos de arrays que podem ou não ter sido definidos previamente
	//Eles possuem dimensão dinâmica

	var outrosNumeros []int

	outrosNumeros = numeros[1:] // Elemento 1 (inclusive) até o final
	outrosNumeros = numeros [1:3] // Elemento 1 (inclusive) até 3 (exclusive)

	fmt.Println(outrosNumeros)

	//O slice não precisa ser necessariamente uma parte de um array ja existente

	nomes := []string{"Ana", "Pedro"}

	fmt.Println(nomes)

	//Adicionar elementos no slice

	nomes = append(nomes, "João")

	fmt.Println(nomes)

	//Iterando sobre arrays ou slides
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

	for indice, num := range numeros{
		fmt.Println("Índice", indice, "- valor", num)
	}

	fmt.Println("------------------")
	modificaArray(numeros)
	fmt.Println(numeros)

	outrosNumeros = numeros[1:3]
	modificaSlice(outrosNumeros)
	fmt.Println(outrosNumeros)
	fmt.Println(numeros)

	slice := criarSlice ()




}

func modificaArray(a[4]int){
	a[0] = 999
}

func modificaSlice(s[]int)