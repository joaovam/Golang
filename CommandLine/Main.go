package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	key1 := flag.String("k1", "", "First key to compare")
	var key2 *string = flag.String("k2", "", "Second key to compare")
	help := flag.Bool("h", false, "Print help.")
	flag.Parse()

	if *help {
		fmt.Println("-key1 \tFirst key to compare")
		fmt.Println("-key2 \tSecond key to compare")
		fmt.Println("-h    \tIf need help")

	}else{
		*key1 = strings.ToLower(*key1)
		*key2 = strings.ToLower(*key2)

		if len(*key1) > 0 && len(*key2) > 0 {
			if *key1 == *key2 {
				fmt.Println("Both keys are equal")

			} else {
				fmt.Println("Strings are different")
			}
		} else {
			fmt.Println("We need two keys to compare. use -h for help")
		}
	}

}
