package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	groups := readInput("../input.txt")
	total := 0
	for _, group := range groups {
		persons := strings.Split(group, "\n")
		var ques []string
		for _, char := range group {
			if !unicode.IsSpace(char) && !stringInSlice(string(char), ques) {
				ques = append(ques, string(char))
			}
		}

		qCount := 0
		for _, que := range ques {
			addIt := true
			for _, ans := range persons {
				if !strings.Contains(ans, que) {
					addIt = false
				}
			}
			if addIt {
				qCount++
			}
		}

		total += qCount
	}

	fmt.Print("Answer: ")
	fmt.Println(total)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Print("Execution time: ")
	fmt.Println(elapsed)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func readInput(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	text := string(content)
	lines := strings.Split(text, "\n\n")

	return lines
}
