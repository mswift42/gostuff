package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queryboard [256][256]int

func main() {

	// Open your file using the second command line arguement becuase the binary name is the first.
	mb := makeBoard()
	file, err := os.Open(os.Args[1])
	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Split(strings.Trim(str, "\n"), " ")
			processQuery(parameters, &mb)

			str, err = reader.ReadString('\n')

		}
	} else {

	}
}

func makeBoard() queryboard {
	newboard := queryboard{}
	return newboard
}
func processQuery(s []string, q *queryboard) {
	switch {
	case s[0] == "SetCol":
		col, _ := strconv.Atoi(s[1])
		val, _ := strconv.Atoi(s[2])
		setCol(q, col, val)
	case s[0] == "SetRow":
		row, _ := strconv.Atoi(s[1])
		val, _ := strconv.Atoi(s[2])
		setRow(q, row, val)
	case s[0] == "QueryCol":
		col, _ := strconv.Atoi(s[1])
		fmt.Println(queryCol(q, col))
	case s[0] == "QueryRow":
		row, _ := strconv.Atoi(s[1])
		fmt.Println(queryRow(q, row))
	}
}

func setCol(q *queryboard, col, val int) *queryboard {
	for i := range q {
		q[i][col] = val
	}
	return q
}
func setRow(q *queryboard, row, val int) *queryboard {
	for i := range q[row] {
		q[row][i] = val
	}
	return q
}
func queryCol(q *queryboard, col int) (sum int) {
	for i := range q {
		sum += q[i][col]
	}
	return sum
}
func queryRow(q *queryboard, row int) (sum int) {
	for i := range q[row] {
		sum += q[row][i]
	}
	return sum
}
