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
			parameters := strings.TrimSpace(str)
			fmt.Println(subSum(digitToInt(parameters)))
			str, err = reader.ReadString('\n')
		}

	} else {
		// You can catch file propblems if you want. It is proper... :)
	}

}
func digitToInt(s string) []int {
	split := strings.Split(strings.Replace(s, " ", "", -1), ",")
	intslice := make([]int, len(split))
	for i, j := range split {
		num, _ := strconv.Atoi(string(j))
		intslice[i] = num
	}
	return intslice
}
func sumList(ints []int) (sum int) {
	for _, i := range ints {
		sum += i
	}
	return sum
}

func subSum(ints []int) int {
	max := sumList(ints)
	sum := 0
	length := len(ints)
	for i := 0; i < length-1; i++ {
		for j := length - 1; j > 0; j-- {
			if j > i {
				sum = sumList(ints[i:j])
				if sum > max {
					max = sum
				}
			}
		}
	}
	return max
}
