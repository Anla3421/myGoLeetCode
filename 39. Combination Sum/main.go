package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(candidates, target)
	fmt.Println("result", result)
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	if len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	temp := []int{}
	helper(candidates, temp, target, &result, 0)
	return result
}

func helper(candidates []int, temp []int, target int, result *[][]int, p int) {
	if p == len(candidates) || target < 0 {
		return
	}

	if target == 0 {
		r := append([]int{}, temp...)
		*result = append(*result, r)
		return
	}

	for i := p; i < len(candidates); i++ {
		t := append(temp, candidates[i])
		helper(candidates, t, target-candidates[i], result, i)
	}
}
