// Same as Part 1, except the Crab turns out to be crazy. Now there is 1,000,000 cups and 10,000,000 moves.
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	input := "643719258"
	var cups []int
	nextCup := make(map[int]int)
	for _, num := range input {
		cups = append(cups, stringToInt(string(num)))
	}
	// Append cups 10-1,000,000 in order
	for i := 10; i < 1000001; i++ {
		cups = append(cups, i)
	}
	// Set up map from each cup to its next cup.
	numOfCups := len(cups)
	for i, cup := range cups {
		if cup == numOfCups {
			nextCup[cup] = cups[0]
		} else {
			nextCup[cup] = cups[i+1]
		}
	}
	currentCup := cups[0]

	for i := 0; i < 10000000; i++ {
		nextCup = playRound(currentCup, numOfCups, nextCup)
		currentCup = nextCup[currentCup]
	}

	num1 := nextCup[1]
	num2 := nextCup[nextCup[1]]

	fmt.Print("Answer: ")
	fmt.Println(num1 * num2)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Print("Execution time: ")
	fmt.Println(elapsed)
}

func playRound(currentCup, numOfCups int, nextCup map[int]int) map[int]int {
	// Find three cups to move + destination cup.
	cup1 := nextCup[currentCup]
	cup2 := nextCup[cup1]
	cup3 := nextCup[cup2]
	dest := currentCup - 1
	if dest == 0 {
		dest = numOfCups
	}
	for dest == cup1 || dest == cup2 || dest == cup3 {
		dest--
		if dest == 0 {
			dest = numOfCups
		}
	}
	// Set new nextCup values.
	nextCup[currentCup] = nextCup[cup3]
	nextCup[cup3] = nextCup[dest]
	nextCup[dest] = cup1
	return nextCup
}

func stringToInt(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return num
}
