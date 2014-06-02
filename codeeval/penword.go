package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Open your file using the second command line arguement becuase the binary name is the first.

	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Split(strings.Trim(str, "\n"), " ")
			fmt.Println(lastWord(parameters))
			str, err = reader.ReadString('\n')

		}
	} else {

	}
}
func lastWord(s []string) (word string) {
	length := len(s)
	return s[length-2]
}
