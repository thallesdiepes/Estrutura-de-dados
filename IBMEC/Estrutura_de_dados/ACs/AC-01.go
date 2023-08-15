package main

import "fmt"

func ePrimo(n int) bool {
    if n <= 1 {
        fmt.Println("O número deve ser maior do que 1")
        return false
    }

    divisores := []int{}

    for i := 2; i <= n/2; i++ {
        if n%i == 0 {
            divisores = append(divisores, i)
        }
    }

    if len(divisores) == 0 {
        return true
    }

    fmt.Println("O número não é primo e é divisível por:", divisores)
    return false
}

func Fibo(n int) int {
	if n <= 0 {
		fmt.Println("O número deve ser positivo")
		return -1
	}
	if n == 1 || n == 2 {
		return 1
	}

	termoAnterior := 1
	termoAtual := 1

	for i := 3; i <= n; i++ {
		proxTermo := termoAnterior + termoAtual
		termoAnterior = termoAtual
		termoAtual = proxTermo
	}

	return termoAtual
}

func diaDaSemana(n int) {
	switch n {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	case 4:
		fmt.Println("Quarta-feira")
	case 5:
		fmt.Println("Quinta-feira")
	case 6:
		fmt.Println("Sexta-feira")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("valor inválido")
	}
}

func eBissexto(ano int) bool {
	if ano%4 != 0 {
		return false
	} else if ano%100 != 0 {
		return true
	} else if ano%400 == 0 {
		return true
	}
	return false
}

func main() {
    fmt.Println("Verificando se o número é primo:")
    num := 12
    ePrimo(num)

    fmt.Println("\nCalculando o n-ésimo número na série de Fibonacci:")
    n := 6
    resultado := Fibo(n)
    fmt.Printf("O %dº número na série de Fibonacci é %d\n", n, resultado)

    fmt.Println("\nExibindo o dia correspondente da semana:")
    dia := 4
    diaDaSemana(dia)

    fmt.Println("\nVerificando se o ano é bissexto:")
    ano := 2000
    if eBissexto(ano) {
        fmt.Println(ano, "é bissexto")
    } else {
        fmt.Println(ano, "não é bissexto")
    }
}

