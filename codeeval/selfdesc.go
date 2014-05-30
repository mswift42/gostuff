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
			fmt.Println(selfDesc(parameters[0]))

			str, err = reader.ReadString('\n')
		}
	} else {

	}
}
func selfDesc(s string) string {
	for i, j := range s {
		c := fmt.Sprintf("%d", i)
		dig, _ := strconv.Atoi(string(j))
		if countChar(c, s) != dig {
			return "0"

		}
	}
	return "1"
}
func countChar(digit, s string) (count int) {
	for _, i := range s {
		if digit == string(i) {
			count++
		}
	}
	return count
}
