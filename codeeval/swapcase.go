package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	// Open your file using the second command line arguement becuase the binary name is the first.

	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Trim(str, "\n")
			fmt.Println(swapCase(parameters))
			//			fmt.Println(swapCase(parameters[0]))

			str, err = reader.ReadString('\n')

		}
	} else {

	}
}

func swapCase(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		} else {
			return r
		}
	}, s)
}
