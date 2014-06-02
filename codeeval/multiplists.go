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
			split := strings.Split(parameters, "|")
			for i := range split {
				split[i] = strings.TrimSpace(split[i])

			}
			spl1 := strings.Split(split[0], " ")
			spl2 := strings.Split(split[1], " ")
			num1 := DigitStringToIntList(spl1)
			num2 := DigitStringToIntList(spl2)
			multiList(num1, num2)
			str, err = reader.ReadString('\n')

		}
	} else {
		// You can catch file propblems if you want. It is
	}
}

func multiList(num1, num2 []int) {
	result := ""
	for i := range num1 {
		num1[i] *= num2[i]
		result += fmt.Sprintf("%d ", num1[i])
	}

	fmt.Println(strings.TrimSpace(result))
}
func DigitStringToIntList(s []string) []int {
	intslice := make([]int, 0)
	for _, j := range s {
		if string(j) != " " {
			num, err := strconv.Atoi(string(j))
			if err != nil {
				fmt.Println("Conversion of string to int failed: ", err)
			}
			intslice = append(intslice, num)
		}
	}
	return intslice
}
