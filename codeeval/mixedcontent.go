package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	file, err := os.Open(os.Args[1])

	if err == nil {

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')

		for err == nil {
			parameters := strings.Trim(str, "\n")
			words := regexp.MustCompile("[a-zA-z]+")
			nums := regexp.MustCompile(`(\d+)`)
			fmt.Println(procLine(parameters, words, nums))

			str, err = reader.ReadString('\n')
		}

	} else {
		// You can catch file propblems if you want. It is
	}
}
func procLine(s string, rew, ren *regexp.Regexp) string {
	if notMixed(s, rew, ren) {
		return s
	}
	return words(s, rew) + "|" + nums(s, ren)
}

func notMixed(s string, rew, ren *regexp.Regexp) bool {
	return !rew.MatchString(s) && ren.MatchString(s)
}
func words(s string, re *regexp.Regexp) string {
	return strings.Join(re.FindAllString(s, -1), ",")
}
func nums(s string, re *regexp.Regexp) string {
	return strings.Join(re.FindAllString(s, -1), ",")
}
