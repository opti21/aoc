package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

