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
			hexToDec(parameters[0])
			str, err = reader.ReadString('\n')
		}
	} else {

	}
}
func hexToDec(s string) {
	hex := "0x" + s
	dec, _ := strconv.ParseInt(hex, 0, 16)
	fmt.Println(dec)
}
