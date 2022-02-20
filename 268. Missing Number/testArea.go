package main

func missingNumber(nums []int) int {
	r := len(nums)
	for i, num := range nums {
		r += i - num
	}
	return r
}
