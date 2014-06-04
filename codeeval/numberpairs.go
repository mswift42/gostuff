package main

import (
	"bufio"
	"fmt"
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
			parameters := strings.Trim(str, "\n")
			if len(parameters) > 0 {
				split := strings.Split(parameters, ";")
				nums, target := split[0], split[1]
				numberPair(nums, target)
			}
			str, err = reader.ReadString('\n')

		}
	} else {

	}
}
func numberPair(s, t string) {
	split := strings.Split(s, ",")
	intslice := make([]int, len(split))
	sum, _ := strconv.Atoi(t)
	result := ""
	for i := range split {
		num, _ := strconv.Atoi(string(split[i]))

		intslice[i] = num

	}
	for _, i := range intslice {
		for _, j := range intslice {
			if i+j == sum && !strings.Contains(result, fmt.Sprintf("%d,%d", j, i)) {
				result += fmt.Sprintf("%d,%d;", i, j)
			}
		}
	}
	if result == "" {
		fmt.Println("NULL")
	}
	fmt.Println(strings.Trim(result, ";"))

}
