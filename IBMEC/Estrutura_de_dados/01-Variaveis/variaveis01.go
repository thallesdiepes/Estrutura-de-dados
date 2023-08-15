package main

import "fmt"

func main() {
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

/* TIPOS DE DADOS

>Inteiros<
	Int8   -128 s 127
	Int16   -32... até 32...
	Int32
	Int64
	Uint8   0 até 255
	Uint16
	Uint32
	Uint64

	Int   32bits ou 64 bits (depende da arquitetura do computador - normalmente 64)
	Byte   Equvalente a uint8
	Rune   Equivalente a int32

------------------

>Flutuantes<

float32
float64 -> Declaração implícita ou curta

------------------

>Complexos<

complex64
complex128

------------------

>Demais<

string
bool -> true / false
nil

*/
