package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	result := containsNearbyDuplicate(nums, k)
	fmt.Println("result", result)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	temp := make(map[int]int, len(nums))
	for index, value := range nums {
		if _, ok := temp[value]; ok {
			if index-temp[value] <= k {
				return true
			}
		}
		fmt.Println(temp)
	}
	return false
}
