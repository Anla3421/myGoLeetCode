package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 1, 5, 5}
	k := 0
	output := findPairs(nums, k)
	fmt.Print(output)
}

func findPairs(nums []int, k int) int {
	if k == 0 {
		output := forZeroK(nums)
		return output
	}
	newNums := removeRepeat(nums)
	var count int
	lens := len(newNums)
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ {
			diff := newNums[i] - newNums[j]
			if diff == k || diff == -k {
				count += 1
			}
		}
	}
	return count
}

// Remove repeated values
func removeRepeat(nums []int) []int {
	newNums := make([]int, 0, len(nums))
	temp := make(map[int]int)
	for _, item := range nums {
		if _, ok := temp[item]; !ok {
			temp[item] = 0
			newNums = append(newNums, item)
		}
	}
	return newNums
}

// for k=0
func forZeroK(nums []int) int {
	var count int
	temp := make(map[int]int)
	for _, item := range nums {
		if _, ok := temp[item]; !ok {
			temp[item] = 0
		}
		temp[item] += 1
	}
	for _, v := range temp {
		if v > 1 {
			count += 1
		}
	}
	fmt.Println(temp)
	return count
}
