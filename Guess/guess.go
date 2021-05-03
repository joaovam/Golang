package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100)

	fmt.Println(" I have chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make a guess: ")

	for i := 0; i < 4; i++ {

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess == target {
			i = 5
			fmt.Println("Correct! Congratulations!")

		} else if guess < target {

			fmt.Println("Target is bigger")
		} else {
			fmt.Println("Target is smaller ")
		}
	}
	fmt.Println("Thanks for playing!")

}
