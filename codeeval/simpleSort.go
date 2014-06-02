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

	// Open your file using the second command line arguement becuase the binary name is the first.

	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Split(strings.Trim(str, "\n"), " ")
			slice := stringToInt(parameters)
			sort.Sort(sort.Float64Slice(slice))
			for _, i := range slice {
				fmt.Printf("%.3f ", i)
			}
			fmt.Println()

			str, err = reader.ReadString('\n')

		}
	} else {

	}
}

func stringToInt(s []string) []float64 {
	slice := []float64{}
	for _, i := range s {
		x, _ := strconv.ParseFloat(string(i), 64)
		slice = append(slice, x)
	}
	return slice
}
