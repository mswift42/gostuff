package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open(os.Args[1])
	stat, _ := file.Stat()
	fmt.Println(stat.Size())

}
