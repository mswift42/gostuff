package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	// Open your file using the second command line arguement becuase the binary name is the first.

	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {

			parameters := strings.TrimSpace(str)
			if len(parameters) > 0 {
				split := strings.Split(parameters, "|")
				findWriter(split[0], split[1])

			}

			str, err = reader.ReadString('\n')

		}
	} else {

	}
}
