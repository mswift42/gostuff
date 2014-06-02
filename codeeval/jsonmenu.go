package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type jsonmenu struct {
	menu struct {
		header string
		items  []struct {
			id    string
			label string
		}
	}
}

func main() {

	// Open your file using the second command line arguement becuase the binary name is the first.

	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {

			parameters := strings.TrimSpace(str)
			if len(parameters) > 0 {
				byt := []byte(parameters)
				var dat map[string]interface{}
				if err := json.Unmarshal(byt, &dat); err != nil {
					panic(err)
				}
				fmt.Println(dat["menu"])
				for key, value := range dat["items"].([]interface{}) {
					fmt.Printf("%T%v%T%v\n", key, key, value, value)
				}

			}
			str, err = reader.ReadString('\n')

		}
	} else {

	}
}
