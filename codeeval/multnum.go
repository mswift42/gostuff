package main

/*
  Here is a sample construct that reads a file name supplied as an argument and
  allows you, the programmer, to easily work with the supplied parameters. Have Fun!
*/

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

	// Anticipate errors
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Split(strings.Split(strings.Trim(str, "\n"), " ")[0], ",")

			a, _ := strconv.Atoi(parameters[0])
			b, _ := strconv.Atoi(parameters[1])
			multNum(a, b)
			str, err = reader.ReadString('\n')
		}
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}
}

func multNum(a, b int) {
	for i := b; ; i = i + b {
		if i >= a {
			fmt.Println(i)
			break
		}
	}
}
