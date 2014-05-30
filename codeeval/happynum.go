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
			parameters := strings.Split(strings.Trim(str, "\n"), ",")
			num, _ := strconv.Atoi(parameters[0])
			fmt.Println(happyNum(num))

			str, err = reader.ReadString('\n')
		}
	} else {

	}
}
func squareDigits(s string) (sum int) {
	for _, i := range s {
		a, _ := strconv.Atoi(string(i))
		sum += a * a
	}
	return sum
}

func happyNum(i int) string {
	num := i
	for {
		num = squareDigits(fmt.Sprintf("%d", num))
		if num == 1 {
			return "1"

		}
		if num == 89 {
			return "0"

		}
	}
	return "-1"
}
