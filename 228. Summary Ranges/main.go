package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{}
	result := summaryRanges(nums)
	fmt.Println("result", result)
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{fmt.Sprintf("%v", nums[0])}
	}
	temp := map[int]int{}
	result := []int{}
	output := []string{}
	keys := []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 != nums[i] {
			temp[nums[i-1]]++
		} else {
			temp[nums[i-1]] = 0
		}
		if i == len(nums)-1 {
			temp[nums[i]]++
		}
	}
	// fmt.Println(temp)
	for k := range temp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		// fmt.Println("k", k)
		result = append(result, k)
		if temp[k] != 0 {
			sort.Ints(result)
			if len(result) < 2 {
				output = append(output, fmt.Sprintf("%v", result[0]))
			} else {
				output = append(output, fmt.Sprintf("%v->%v", result[0], result[len(result)-1]))
			}
			result = nil
		}
	}
	return output
}
