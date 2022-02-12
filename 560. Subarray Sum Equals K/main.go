package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 2, 1}
	// nums := []int{1, -1, 0}
	k := 3
	result := subarraySum(nums, k)
	fmt.Println("result", result)
}

func subarraySum(nums []int, k int) int {
	// the hashmap prefix stores the prefix sum from 0 to i and the number of its occurence.
	// prefix := map[int]int{0: 1}
	prefix := map[int]int{0: 1}
	var sum int
	var count int
	fmt.Println("prefix", prefix)
	for _, num := range nums {
		// sum represents the sum of nums[:i](including nums[i])
		sum += num
		fmt.Println("sum", sum)
		// since the hashmap prefix has stored the sum of the subarray nums[:j](0<=j<=i) and
		// the number of its occurence, so the number of the subarray nums[j:i](including i
		// and j) whose sum is k is equal to the number of the prefix nums[:j] whose sum is
		// sum-k.
		count += prefix[sum-k]
		fmt.Println("sum-k", sum-k)
		fmt.Println("count", count)
		fmt.Println("prefix[sum-k]", prefix[sum-k])
		fmt.Println("prefix[sum]", prefix[sum])
		prefix[sum]++
		fmt.Println("prefix[sum]++", prefix[sum])
		fmt.Println()
	}

	return count
}

// func subarraySum(nums []int, k int) int {
// 	// aaa := Sum(nums, 0, 0) // No.
// 	var aaa int
// 	var equals []int
// 	var output int
// 	lens := len(nums)
// 	for _, v := range nums {
// 		if v == k {
// 			output += 1
// 		}
// 	}
// 	for f := lens; f > -1; f-- {
// 		for i := lens - 1; i > -1; i-- {
// 			result, sumF, sumI := Sum(nums, f, i)
// 			aaa = sumF + sumI
// 			if result == k {
// 				temp, err := fmt.Printf("%d%d ", f, i)
// 				if err != nil {
// 					fmt.Println(err)
// 					panic(0)
// 				}
// 				equals = append(equals, temp)
// 			}
// 		}
// 	}
// 	fmt.Println("sum", aaa)
// 	output += len(equals)
// 	return output
// }

// func Sum(nums []int, final int, init int) (int, int, int) {
// 	var (
// 		output int
// 		sumI   int
// 		sumF   int
// 	)
// 	for i := 0; i < init; i++ {
// 		sumI += nums[i]
// 	}
// 	for f := 0; f < final; f++ {
// 		sumF += nums[f]
// 	}
// 	output = sumF - sumI
// 	return output, sumF, sumI
// }
