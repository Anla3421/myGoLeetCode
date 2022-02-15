package main

import (
	"fmt"
)

func main() {
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := intersect(nums1, nums2)
	fmt.Println("result", result)
}

func intersect(nums1 []int, nums2 []int) []int {
	nums1Temp := map[int]int{}
	result := []int{}
	for _, v := range nums1 {
		nums1Temp[v]++
	}
	for _, v := range nums2 {
		if _, ok := nums1Temp[v]; ok {
			if nums1Temp[v] > 0 {
				nums1Temp[v]--
				result = append(result, v)
				if nums1Temp[v] == 0 {
					delete(nums1Temp, v)
				}
			}
		}
	}
	return result
}
