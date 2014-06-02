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
			parameters := strings.Trim(str, "\n")
			calcDist(parameters)

			//			fmt.Println(swapCase(parameters[0]))

			str, err = reader.ReadString('\n')

		}
	} else {

	}
}
func digtoint(s []string) []int {
	intslice := make([]int, len(s))
	for i, j := range s {
		n, _ := strconv.Atoi(string(j))
		intslice[i] = n
	}
	return intslice
}

func calcDist(s string) {
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	s = strings.Replace(s, ",", "", -1)
	split := strings.Split(s, " ")
	intslice := digtoint(split)
	a1, b1, a2, b2 := intslice[0], intslice[1], intslice[2], intslice[3]
	fmt.Println(int(math.Sqrt(math.Pow(float64(a2-a1), 2) + math.Pow(float64(b2-b1), 2))))

}
