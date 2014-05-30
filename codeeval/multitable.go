package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i < 13; i++ {
		line := ""
		for a, b := i, 1; b < 13; b++ {
			line += fmt.Sprintf("%4d", a*b)
		}
		fmt.Println(strings.Trim(line, " "))

	}

}
