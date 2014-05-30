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
			line := strings.Split(parameters[0], ",")
			num, pos1, pos2 := line[0], line[1], line[2]
			checkPos(num, pos1, pos2)

			str, err = reader.ReadString('\n')
		}
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}
}

func checkPos(num, pos1, pos2 string) {
	n, _ := strconv.Atoi(num)
	p1, _ := strconv.Atoi(pos1)
	p2, _ := strconv.Atoi(pos2)

	number := int64(n)
	numdigits := strconv.FormatInt(number, 2)
	length := len(numdigits)

	if numdigits[length-p1] == numdigits[length-p2] {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
