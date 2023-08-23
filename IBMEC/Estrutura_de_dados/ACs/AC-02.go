package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contato struct {
	Nome  string
	Email string
}

func main() {
	var contatos [5]Contato

	contatos[0] = Contato{Nome: "Thalles"}

	if (contatos[0] == Contato{}) {
		fmt.Println("Elemento 0 está vazio!")
	} else {
		fmt.Println("Elemento 0 não está vazio!")
	}

	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("Informe o seu nome:")
	msg, _ := leitor.ReadString('\n') // '\n' é um byte
	fmt.Println(msg)
}

func adicionaContato(nome string, email string, a [5]Contato) [5]Contato {
	a[0] = Contato{Nome: nome, Email: email}
	return a
}
