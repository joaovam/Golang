// +build !linux

package main

import (
	"fmt"
	"strings"
)

const PathSeparator string = "\\"

func Join(parts ... string) string{
	return strings.Join(parts,PathSeparator)

}

func main() {
	s := Join("a", "b")
	fmt.Println(s)
}
