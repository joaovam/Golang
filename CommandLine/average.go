package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/*Using command-line arguments*/
func main(){

	arguments := os.Args[1:]
	if(len(arguments)==0){
		log.Fatal("Please pass values as arguments")
	}

	var sum float64 = 0

	for _,value := range(arguments){


		number,err := strconv.ParseFloat(value,64)

		if err!=nil{
			log.Fatal(err)
		}
		sum+=number
	}

	fmt.Println("Average =" , sum/float64(len(arguments)))



}