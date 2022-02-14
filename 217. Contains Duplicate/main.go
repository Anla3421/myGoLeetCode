package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result := containsDuplicate(nums)
	fmt.Println("result", result)
}

func containsDuplicate(nums []int) bool {
	temp := map[int]int{}
	for _, v := range nums {
		temp[v]++
		if temp[v] > 1 {
			return true
		}
	}
	return false
}
