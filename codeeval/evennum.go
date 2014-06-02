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
			parameters := strings.Trim(str, "\n")
			fmt.Println(checkEven(parameters))

			str, err = reader.ReadString('\n')

		}
	} else {

	}
}

func checkEven(s string) string {
	num, _ := strconv.Atoi(s)
	if num%2 == 0 {
		return "1"
	}
	return "0"
}
