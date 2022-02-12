package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3}
	target := 6
	result := twoSum(nums, target)
	fmt.Println("result", result)
}

func twoSum(nums []int, target int) []int {
	mapNums := map[int]int{}
	for index, value := range nums {
		if k, ok := mapNums[target-value]; ok {
			return []int{k, index}
		}
		mapNums[value] = index
	}
	return []int{}
}
