package main

import (
	"fmt"
	Carro "testes/Struct/carro"
)



func main(){
	

	carro:= new(Carro.Carro)
	carro.Ano = 2005
	carro.Marca = "Fiat"
	carro.Modelo = "Palio"
	carro.Preco = 10000.500

	carro2 := new(Carro.Carro)	
	carro2.Ano = 2020
	carro2.Marca = "Toyota"
	carro2.Modelo = "Hilux"
	carro2.Preco = 99800.800
	
	fmt.Println("Carro")
	carro.Write()
	fmt.Println("Carro2")
	carro2.Write()

	carro2 = carro.Swap(carro2)

	fmt.Println("Carro")
	carro.Write()
	fmt.Println("Carro2")
	carro2.Write()
	
}

