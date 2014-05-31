package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
			joined := strings.Replace(strings.Map(onlyLowerAlpha, parameters), " ", "", -1)
			m := mapString(joined)
			fmt.Println(scoreMap(m))
			str, err = reader.ReadString('\n')
		}
	} else {

	}
}
func onlyLowerAlpha(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return r + 32
	case r >= 'a' && r <= 'z':
		return r

	}
	return ' '
}
func scoreMap(sm map[string]int) (sum int) {
	sorted := sortedKeys(sm)
	length := len(sorted)
	for i, j := 0, 26; i < length; i, j = i+1, j-1 {
		sum += sm[sorted[i]] * j
	}
	return sum
}

func mapString(s string) map[string]int {
	m := make(map[string]int)
	for _, i := range s {
		x := string(i)
		_, ok := m[x]
		if !ok {
			m[x] = 1
		} else {
			m[x]++
		}
	}
	return m
}

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}
