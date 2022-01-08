package main

import (
	"fmt"

	"github.com/opti21/AOC/utils"
)

func main() {
	input := utils.InputFile("./input.txt")

	resultA := make(chan int64)
	resultB := make(chan int64)

	go puzzleA(input, resultA)
	go puzzleB(input, resultB)

	fmt.Println(<-resultA)
	fmt.Println(<-resultB)
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

	result <- utils.BinaryStringToInt(gamma) * utils.BinaryStringToInt(epsilon)

}

func puzzleB(input []string, result chan int64) {
	ones := [12]int{0}
	zeros := [12]int{0}
	count := 0

	for _, line := range input {
		count += 1
		for i, digit := range line {
			switch digit {
			case '1':
				ones[i] += 1
			case '0':
				zeros[i] += 1
			}
		}
	}


	oxygenBinary := ""
	// cO2Binary := ""


	oxygenList := input
	// cO2List := input

	for inputIdx := 0; inputIdx < len(input[0]); inputIdx++ {
		for len(oxygenList) > 1 {
				newOxyList := []string{}
				for _, line := range oxygenList {
					if  line[inputIdx] == 0 {

						newOxyList = append(newOxyList, line)

					}
				}

				oxygenList = newOxyList

		}


	}


	// for len(cO2List) > 1 {
	// 	for binaryIdx := 0; binaryIdx < len(cO2Binary) - 1; binaryIdx++ {
	// 		newCO2List := []string{}
	// 		for _, line := range cO2List {
	// 			if cO2Binary[binaryIdx] == line[binaryIdx] {
	// 				newCO2List = append(newCO2List, line)
	// 			}
	// 		}

	// 		cO2List = newCO2List

	// 	}

	// }


	result <- 0

}