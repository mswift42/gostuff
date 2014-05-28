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
			parameters := strings.Split(strings.Trim(str, "\n"), ",")
			trailingString(parameters[0], parameters[1])

			str, err = reader.ReadString('\n')
		}
	} else {

	}
}
func trailingString(a, b string) {
	split := strings.Split(a, " ")
	last := len(split) - 1
	endofa := split[last]
	if endofa == b {
		fmt.Println("1")
	} else {
		fmt.Print("0")
	}
}
