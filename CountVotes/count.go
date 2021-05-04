package main

import (
	"fmt"
	"log"
	datafile "testes/DataFile"
)

//Count votes from a file, tests the usage of a Map
func main(){

	lines,err := datafile.GetStrings("votes.txt")

	if err!=nil{
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _,vote := range(lines){

		counts[vote]++
	}

	PrintVotes(counts)

}


func PrintVotes(maps map[string]int){
	max := 0 
	for _,vote := range (maps){
		max+=vote
		
	}
	for name,vote := range (maps){

		fmt.Printf("%v has a grade of %v %% \n",name,(float64(vote)/float64(max)*100))

	}

}