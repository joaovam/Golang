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

func (e Carro) Write() {
	fmt.Printf("[Modelo: %v, Ano: %v, preco: %v, marca: %v ]\n",e.Modelo,e.Ano,e.Preco,e.Marca)
}
