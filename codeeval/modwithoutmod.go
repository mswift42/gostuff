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
			parameters := strings.Split(strings.Trim(str, "\n"), ",")
			mod(parameters[0], parameters[1])
			str, err = reader.ReadString('\n')
		}
	} else {

	}
}

func mod(a, b string) {
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	c := n / m
	fmt.Println(n - m*c)
}
