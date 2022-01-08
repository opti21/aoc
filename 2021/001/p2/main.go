package main

import (
	"fmt"

	"github.com/opti21/AOC/utils"
)


func main() {
	depth_data := utils.ReadIntsFromInput("../input.txt")
	depths_increased := 0

	for i, depth := range depth_data {
		var prev int
		var next int

		if i != len(depth_data) - 3 {
			next = depth_data[i + 1] + depth_data[i + 2] + depth_data[i + 3]
			prev = depth + depth_data[i + 1] + depth_data[i + 2]
		} else {
			break
		}

		if next > prev {
			depths_increased++
		}
	}

	fmt.Println(depths_increased)
}