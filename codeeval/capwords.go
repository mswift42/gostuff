package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {

	// Open your file using the second command line arguement becuase the binary name is the first.

	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Split(strings.Trim(str, "\n"), " ")
			fmt.Println(capWords(parameters))
			str, err = reader.ReadString('\n')

		}
	} else {

	}
}

func capWords(s []string) (sentence string) {
	for _, i := range s {
		r, n := utf8.DecodeRuneInString(i)
		sentence += string(unicode.ToUpper(r)) + i[n:]
		sentence += " "
	}
	return strings.Trim(sentence, " ")
}
