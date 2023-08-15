package main

import "fmt"

type Pessoa struct {
	Nome	string //isso é um campo
	Idade	int //mais um campo
	Altura	float64 //outro campo
}

func (p *Pessoa) AvancarIdade() {
	p.Idade++
}

type Ponto struct{
	x, y	int
}

type Retangulo struct{
	Ponto //Campo anônimo
	Largura, Altura	int
}

type Circulo struct{
	Raio	float64
}

func (c Circulo) Area() float64{
	return math.Pi * c.Raio * c.Raio
}

func main()

func main() {
	p := Pessoa{
		Nome:	"Alice",
		Idade:	25,
		Altura:	1,65,
	}

	fmt.Println(p)
	fmt.Println(p.Nome)
	fmt.Println(p.Altura)

	r := Retangulo{
		Ponto:	Ponto{}
	}

}