package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			mthToLast(parameters)

			str, err = reader.ReadString('\n')
		}
	} else {

	}
}
func mthToLast(s []string) {
	length := len(s)
	index, _ := strconv.Atoi(s[length-1])
	if index > length {
		fmt.Println()

	} else {
		fmt.Println(s[length-1-index])
	}
}
