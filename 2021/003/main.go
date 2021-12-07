package main

import (
	"fmt"

	"github.com/opti21/AOC/utils"
)

func main() {
	input := utils.InputFile("./input.txt")

	resultA := make(chan int64)
	// resultB := make(chan int)

	go puzzleA(input, resultA)
	// go puzzleB(input, resultB)

	fmt.Println(<-resultA)
	// fmt.Println(<-resultB)
}

func puzzleA(input []string, result chan int64) {
	ones := [12]int{0}
	count := 0

	for _, line := range input {
		count += 1
		for i, digit := range line {
			if digit == '1' {
				ones[i] += 1
			}
		}
	}
	gamma := ""
	epsilon := ""

	for _, c := range ones {
		if c > count/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	fmt.Println(ones)
	fmt.Printf("gamma: %v\n", gamma)
	fmt.Printf("epsilon: %v\n", epsilon)

	result <- utils.BinaryStringToInt(gamma) * utils.BinaryStringToInt(epsilon)

}

func puzzleB(input [][]int, result chan int) {
}