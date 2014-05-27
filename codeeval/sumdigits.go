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
			sumDigits(parameters[0])

			// Parameters is a slice of strings that represent the space delimited
			// items located on each line of the file.

			// At this point you would parse your strings if you need a different type ie
			// a := strconv.Atoi(params[0])

			// In my applicaitons I have used a single funciton to perform the proper action
			// and return the anticipated result which is printed ie:
			// fmt.Println(a)

			str, err = reader.ReadString('\n')
		}
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}
}

func sumDigits(s string) {
	spl := strings.Split(s, "")
	result := 0
	for _, i := range spl {
		dig, _ := strconv.Atoi(i)
		result += dig
	}
	fmt.Println(result)
}
