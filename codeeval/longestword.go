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
			parameters := strings.Trim(str, "\n")
			split := strings.Split(parameters, " ")
			fmt.Println(longestWord(split))
			str, err = reader.ReadString('\n')
		}
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}
}
func longestWord(s []string) string {
	longest := longestWordSize(s)
	for _, i := range s {
		if len(i) == longest {
			return i
		}
	}
	return ""
}

func longestWordSize(s []string) int {
	max := 0
	for _, i := range s {
		if len(i) > max {
			max = len(i)
		}
	}
	return max
}
