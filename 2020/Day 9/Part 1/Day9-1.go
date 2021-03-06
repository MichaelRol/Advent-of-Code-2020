package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	numbers := readNumbers("../input.txt")

	for i := 25; i < len(numbers); i++ {
		// After first 25 numbers (the preamble)
		// Checks if each number in list is a sum of two of the previous 25 numbers
		if isValid(numbers[i-25:i], numbers[i]) {
			continue
		}
		fmt.Print("Answer: ")
		fmt.Println(numbers[i])
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Print("Execution time: ")
		fmt.Println(elapsed)
		return
	}
}

func isValid(numbers []int, value int) bool {
	for _, i := range numbers {
		for _, j := range numbers {
			if i == j {
				continue
			} else if i+j == value {
				return true
			}
		}
	}
	return false
}

func readNumbers(filename string) []int {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	text := string(content)
	lines := strings.Split(text, "\n")

	var numbers []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		numbers = append(numbers, i)
	}

	return numbers
}
