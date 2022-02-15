package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 1, 2, 1, 2}
	result := singleNumber(nums)
	fmt.Println("result", result)
}

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		fmt.Println(res)
		res = res ^ n
		fmt.Println(res)
	}
	return res
}
