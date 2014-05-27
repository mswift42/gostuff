package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Split(strings.Trim(str, "\n"), " ")
			pop(parameters)
			fmt.Println()

			str, err = reader.ReadString('\n')
		}
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}
}

func pop(s []string) {
	l := len(s)
	for i := l - 1; i >= 0; i = i - 2 {
		fmt.Printf("%s ", s[i])
	}
}
