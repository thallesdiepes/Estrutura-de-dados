package main

import "fmt"

func main(){
	var x = 10
	var dia = "segunda"

	if x > 18{
		fmt.Println("Você é maior de idade")
	} else if x < 0 {
		fmt.Println("Valor inválido")
	} else {
		fmt.Println("Você é menor de idade")
	}

	switch dia {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("Dia ítil")
	case "sábado", "domingo":
		fmt.Println("Fim de semana")
	default:
		fmt.Println("Dia inválido")
	}
}

