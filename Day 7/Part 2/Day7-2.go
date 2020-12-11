package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	lines := strings.Split(text, "\n")

	rules := make(map[string][][]string)

	for _, line := range lines {
		both := strings.Split(line, " contain ")
		contains := strings.Split(both[1], ", ")
		for _, color := range contains {
			if color[len(color)-1] == '.' {
				color = color[:len(color)-1]
			}
			if color[len(color)-1] == 's' {
				color = color[:len(color)-1]
			}
			var numAndName []string
			if color[:1] == "n" {
				numAndName = append(numAndName, "0", "no other bags")
			} else {
				numAndName = append(numAndName, color[:1], color[2:])
			}

			rules[both[0][:len(both[0])-1]] = append(rules[both[0][:len(both[0])-1]], numAndName)
		}
	}

	var numAndName []string
	numAndName = append(numAndName, "1", "shiny gold bag")
	fmt.Println(countBags(numAndName, rules) - 1)
}

func countBags(bag []string, rules map[string][][]string) int {
	total := 0
	if bag[0] == "0" {
		return 0
	}

	for _, nextBag := range rules[bag[1]] {
		i, err := strconv.Atoi(nextBag[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		total += i * countBags(nextBag, rules)
	}

	return 1 + total
}
