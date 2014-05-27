package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Split(strings.Trim(str, "\n"), ",")
			origin := strings.Split(parameters[0], ",")[0]
			deletechars := strings.Trim(parameters[1], " ")
			remChars(origin, deletechars)
			x := 1
			if *(*byte)(unsafe.Pointer(&x)) == 1 {
				fmt.Println("little endian")
			} else {
				fmt.Println("big endian")
			}
			str, err = reader.ReadString('\n')
		}
	} else {

	}
}

func remChars(origin, delchars string) {
	str := origin
	for _, i := range delchars {
		str = strings.Replace(str, string(i), "", -1)
	}
	fmt.Println(str)

}
