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
				numberPair("2,4,5,6,9,11,15", "20")
			}
			str, err = reader.ReadString('\n')

		}
	} else {

	}
}
func numberPair(s, t string) {
	intslice := make([]int, 0)
	sum, _ := strconv.Atoi(t)
	split := strings.Split(s, "")
	result := ""
	for _, i := range split {
		num, _ := strconv.Atoi(string(i))
		if i != "," {
			intslice = append(intslice, num)
		}
	}
	for _, i := range intslice {
		for _, j := range intslice {
			if i+j == sum && !strings.Contains(result, fmt.Sprintf("%d,%d", j, i)) {
				result += fmt.Sprintf("%d,%d;", i, j)
			}
		}
	}
	fmt.Println(result)

}
