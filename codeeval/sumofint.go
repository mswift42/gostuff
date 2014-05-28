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
		sum := 0
		for err == nil {
			parameters := strings.Split(strings.Trim(str, "\n"), " ")
			num, _ := strconv.Atoi(parameters[0])
			sum += num

			str, err = reader.ReadString('\n')
		}
		fmt.Println(sum)
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}

}
