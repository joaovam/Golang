package datafile

import (
	"bufio"
	"os"
	"strconv"
)

//Receives a string indicating the file name and from this file read a string array
func GetStrings(fileName string)([]string,error){
	file,err := os.Open(fileName)
	var strings []string

	if err==nil {

		scanner :=bufio.NewScanner(file)
		


		for scanner.Scan(){
			strings = append(strings,scanner.Text())
		}
	}else{
		strings = nil
	}
	return strings,err
}

//Receives a string indicating the file name and from this file read a float array 
func GetFloats(fileName string)([]float64,error){
	var numbers []float64

	file,err := os.Open(fileName)

	if err!=nil{

		numbers = nil
	}else{

		scanner := bufio.NewScanner(file)

		for scanner.Scan(){
			data,err := strconv.ParseFloat(scanner.Text(),64)

			if(err!=nil){
				return nil,err
			}

			numbers = append(numbers,data)

		}
	}

	return numbers, err
}