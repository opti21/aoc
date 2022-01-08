package main

import (
	"fmt"

	"github.com/opti21/AOC/utils"
)

func main() {
	input := utils.ReadCommandsFromInput("./input.txt")

	resultA := make(chan int)
	resultB := make(chan int)

	go puzzleA(input, resultA)
	go puzzleB(input, resultB)

	fmt.Println(<-resultA)
	fmt.Println(<-resultB)
}

func puzzleA(input []utils.Command, result chan int) {
	horz := 0
	depth := 0

	for _, command := range input {

		switch command.Direction {
		case "forward":
			horz += command.Amount
		case "down":
			depth += command.Amount
		case "up":
			depth -= command.Amount
		}
	}

	result <- horz * depth
}

func puzzleB(input []utils.Command, result chan int) {
	horz := 0
	depth := 0
	aim := 0

	for _, command := range input {

		switch command.Direction {
		case "forward":
			horz += command.Amount
			depth += aim * command.Amount
		case "down":
			aim += command.Amount
		case "up":
			aim -= command.Amount
		}
	}

	result <- horz * depth

}