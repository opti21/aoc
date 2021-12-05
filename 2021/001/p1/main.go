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
		if i == 0 {
			prev = depth
		} else {
			prev = depth_data[i-1]
		}
		curr := depth

		if curr > prev {
			depths_increased++
		}
	}

	fmt.Println(depths_increased)
}