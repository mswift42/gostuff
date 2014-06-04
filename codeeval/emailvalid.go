package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	distBuild int = 6
)

func main() {

	// Open your file using the second command line arguement becuase the binary name is the first.

	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Trim(str, "\n")

			if len(parameters) > 0 {
				fmt.Println(validEmail(parameters))
				var str string = "Hello"
				unicodeCodePoints := []int(str)
				fmt.Println(unicodeCodePoints)

			}
			str, err = reader.ReadString('\n')

		}
	} else {

	}
}

func validEmail(s string) string {
	re := regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")
	if re.MatchString(s) {
		return "true"
	}
	return "false"
}
