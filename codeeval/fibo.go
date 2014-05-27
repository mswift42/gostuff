package main

import (
	"bufio"
	"fmt"
	"os"

	//"strconv"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open(os.Args[1])

	// Anticipate errors
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Split(strings.Trim(str, "\n"), " ")
			num, _ := strconv.Atoi(parameters[0])
			fib(num)

			str, err = reader.ReadString('\n')
		}
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}
}

func fib(n int) {
	a, b := 1, 1
	for i := int64(1); i < int64(n); i++ {
		a, b = b, a+b
	}
	fmt.Println(a)
}
