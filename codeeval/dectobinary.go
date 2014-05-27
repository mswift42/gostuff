package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Split(strings.Trim(str, "\n"), " ")
			decToBinary(parameters[0])

			str, err = reader.ReadString('\n')
		}
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}
}

func decToBinary(s string) {
	num, _ := strconv.Atoi(s)
	i := int64(num)
	digits := strconv.FormatInt(i, 2)
	fmt.Println(digits)
}
