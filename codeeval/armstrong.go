package main

import (
	"bufio"
	"fmt"
	"math"
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
			armstrong(parameters[0])
			str, err = reader.ReadString('\n')
		}
	} else {

	}
}
func armstrong(s string) {
	num, _ := strconv.Atoi(s)
	sum := 0
	length := float64(len(s))
	for _, j := range s {
		x, _ := strconv.Atoi(string(j))
		sum += int(math.Pow(float64(x), length))
	}
	if num == sum {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
