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
			fmt.Println(substWordWithDigit(parameters))
			str, err = reader.ReadString('\n')
		}
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}
}
func substWordWithDigit(s string) string {
	nummap := map[string]string{"zero": "0", "one": "1",
		"two": "2", "three": "3", "four": "4", "five": "5",
		"six": "6", "seven": "7", "eight": "8", "nine": "9", ";": ""}
	old := s
	for key, value := range nummap {
		old = strings.Replace(old, key, value, -1)
	}
	return old
}
