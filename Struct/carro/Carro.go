package carro

import (
	"fmt"
)

type Carro struct{
	Modelo string
	Ano int
	Preco float64
	Marca string
}

func(e Carro) Carro(){
	e.Ano = 0
	e.Modelo = ""
	e.Preco = 0.0
	e.Marca = ""
}

func (e Carro) Write() {
	fmt.Printf("[Modelo: %v, Ano: %v, preco: %.3f, marca: %v ]\n",e.Modelo,e.Ano,e.Preco,e.Marca)
}


func (e *Carro) Swap(c *Carro) (*Carro){

	aux := *e

	*e = *c

	*c = aux
	return c

}