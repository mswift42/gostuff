package main

/*
  Here is a sample construct that reads a file name supplied as an argument and
  allows you, the programmer, to easily work with the supplied parameters. Have Fun!
*/

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
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
			parameters := strings.Split(strings.Trim(str, "\n"), " ")
			fb(parameters[0], parameters[1], parameters[2])

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

func fb(a, b, c string) {
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	z, _ := strconv.Atoi(c)
	for i := 1; i <= z; i++ {
		if i%x == 0 && i%y == 0 {
			fmt.Printf("FB ")
		} else if i%x == 0 {
			fmt.Printf("F ")
		} else if i%y == 0 {
			fmt.Printf("B ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
