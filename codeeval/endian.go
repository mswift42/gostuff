package main

import (
	"fmt"
	"unsafe"
)

func main() {

	x := 1
	if *(*byte)(unsafe.Pointer(&x)) == 1 {
		fmt.Println("LittleEndian")
	} else {
		fmt.Println("BigEndian")
	}

}
