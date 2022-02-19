package main

import (
	"fmt"
)

func main() {
	nums := []int{6, 6, 6, 7, 7}
	result := majorityElement(nums)
	fmt.Println("result", result)
}

func majorityElement(nums []int) int {
	var (
		maxCounts int
		output    int
	)
	temp := map[int]int{}
	for _, v := range nums {
		temp[v]++
	}
	for k, v := range temp {
		if v > maxCounts {
			output = k
			maxCounts = v
		}

	}
	return output
}
