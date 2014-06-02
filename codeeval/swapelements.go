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
			parameters := strings.Trim(str, "\n")

			split := strings.Split(parameters, ": ")
			numbers := strings.Fields(split[0])
			swaps := strings.Fields(strings.Replace(split[1], ",", "", -1))
			swapDigits(numbers, swaps)
			str, err = reader.ReadString('\n')

		}
	} else {
		// You can catch file propblems if you want. It is
	}
}

func swapDigits(numbers []string, swaps []string) {
	num := make([]int, 0)
	for _, i := range numbers {
		n, _ := strconv.Atoi(i)
		num = append(num, n)
	}
	for _, i := range swaps {
		a, _ := strconv.Atoi(string(i[0]))
		b, _ := strconv.Atoi(string(i[2]))
		temp := num[b]
		num[b] = num[a]
		num[a] = temp
	}
	numstring := ""
	for _, i := range num {
		numstring += fmt.Sprintf("%d ", i)
	}
	fmt.Println(strings.Trim(numstring, " "))

}
