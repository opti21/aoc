package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadIntsFromInput(inputfile string) (result []int) {
	file, err := os.Open(inputfile)

	if err != nil {
		log.Fatalf("failed to open input")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for _, text_line := range text {
		depth, err := strconv.Atoi(text_line)

		if err != nil {
			log.Fatalf("error parsing string")
		}
		result = append(result, depth)
	}

	return

}

type Command struct {
	Direction string
	Amount int
}

func ReadCommandsFromInput (inputfile string) (commands []Command) {
	file, err := os.Open(inputfile)

	if err != nil {
		log.Fatalf("failed to open input")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split_text := strings.Split(scanner.Text(), " ")
		amountInt, err := strconv.Atoi(split_text[1])

		if err != nil {
			log.Fatalf("Error converting string to int")	
		}

		commands = append(commands, Command{split_text[0], amountInt})
	}
	file.Close()


	return

}

func InputToSliceOfSlice (inputfile string) (result [][]int) {
	file, err := os.Open(inputfile)

	if err != nil {
		log.Fatalf("failed to open input")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		slice_text := strings.Split(scanner.Text(), "")

		nums := make([]int, len(slice_text))

		for i := range slice_text {
			nums[i], _ = strconv.Atoi(slice_text[i])
		}

		result = append(result, nums)
	}
	file.Close()

	return

}

func MostFrequent(array []int) int {
	m := map[int]int{}
	var maxCount int
	var freq int
	for _, a := range array {
		m[a]++
		if m[a] > maxCount {
			maxCount = m[a]
			freq = a
		}
	}

	return freq
}

func BinaryStringToInt (input string) int64 {
	output, err := strconv.ParseInt(input, 2, 64)

	if err != nil {
		log.Fatalf("error converting string to int")
	}

	return output
}

func InputFile (inputfile string) (result []string) {
	file, err := os.Open(inputfile)

	if err != nil {
		log.Fatalf("failed to open input")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	file.Close()

	return

}
