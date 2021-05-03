package main

import "fmt"

//	"testes/classes/carro"

func main(){
	
// 	auto := carro.Carro{Modelo:"Taycan",Ano:2020,Preco:909000.900,Marca:"Porsche"}
// 	auto.Write()
// fmt.Println("Finalizado")
	done:= make(chan string)
	//fmt.Println(teste(5))
	fmt.Printf("%T",done)

	
}

func teste(t int) int{
	return t *2;
}