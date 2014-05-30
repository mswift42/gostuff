package main

import (
	"bufio"
	"fmt"
	"os"
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
			fmt.Println(rightMost(parameters[0], parameters[1]))

			str, err = reader.ReadString('\n')
		}
	} else {

	}
}
func rightMost(str, ind string) string {
	length := len(str)
	for i := length - 1; i > 0; i-- {
		if ind == string(str[i]) {
			return fmt.Sprintf("%d", i)
		}
	}
	return "-1"
}
