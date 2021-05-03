package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)




func main(){

	arguments := os.Args
	slice := make([]float64,0)

	for _, value :=range(arguments){

		valueAsFloat,_ := strconv.ParseFloat(value,64)

		slice = append(slice,valueAsFloat)

	}
	fmt.Println(maximum(slice...))

}


func maximum(param ...float64)float64{

	max := math.Inf(-1)


	for _,value := range(param){

		if(max < value){
			max = value
		}

	}
	return max
}