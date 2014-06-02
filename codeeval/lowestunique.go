package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open(os.Args[1])

	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Join(strings.Fields(strings.Trim(str, "\n")), "")
			conv := DigitStringToIntList(parameters)
			fmt.Println(lowestUnique(conv))
			str, err = reader.ReadString('\n')
		}
	} else {
		// You can catch file propblems if you want. It is proper... :)
	}
}
func lowestUnique(xs []int) int {
	min := 0
	single := onlySingleElem(xs)
	sort.Sort(sort.IntSlice(single))
	if len(single) > 0 {
		min = single[0]
		for i := range xs {
			if xs[i] == min {
				return i + 1
			}
		}
	}
	return 0
}

func onlySingleElem(xs []int) []int {
	result := make([]int, 0)
	for _, i := range xs {
		if singleElem(i, xs) {
			result = append(result, i)
		}
	}
	return result
}

func singleElem(n int, xs []int) bool {
	count := 0
	for _, i := range xs {
		if i == n {
			count++
		}
	}
	return count == 1
}
func DigitStringToIntList(s string) []int {
	intslice := make([]int, len(s))
	for i, j := range s {
		num, err := strconv.Atoi(string(j))
		if err != nil {
			fmt.Println("Conversion of string to int failed: ", err)
		}
		intslice[i] = num
	}
	return intslice
}
