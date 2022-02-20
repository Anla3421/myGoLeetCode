package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 0, 1}
	result := missingNumber(nums)
	fmt.Println("result", result)
}

func missingNumber(nums []int) int {
	lens := len(nums)
	var (
		sumofN    int
		sumofNums int
	)

	sumofN = (1 + lens) * lens / 2
	for _, v := range nums {
		sumofNums += v
	}
	return sumofN - sumofNums
}
