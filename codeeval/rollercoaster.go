package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	str, nerr := reader.ReadString('\n')

	for nerr == nil {
		parameters := strings.Trim(str, "\n")
		fmt.Println(rollercoast(parameters))

		str, nerr = reader.ReadString('\n')
	}
}

func rollercoast(s string) string {
	result := ""
	upper := true
	for _, i := range s {
		if !unicode.IsLetter(i) {
			result += string(i)
		} else if upper {
			result += string(unicode.ToUpper(i))
			upper = false
		} else {
			result += string(i)
			upper = true
		}

	}
	return result
}
