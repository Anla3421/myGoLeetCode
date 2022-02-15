package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := intersection(nums1, nums2)
	fmt.Println("result", result)
}

func intersection(nums1 []int, nums2 []int) []int {
	nums1Temp := map[int]struct{}{}
	result := []int{}
	for _, v := range nums1 {
		nums1Temp[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := nums1Temp[v]; ok {
			result = append(result, v)
			delete(nums1Temp, v)
		}
	}
	return result
}
