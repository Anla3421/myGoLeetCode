package main

// failed at[1,2,1,2,1] k=3 ans:4
// func subarraySum(nums []int, k int) int {
// 	var output int
// 	temp := 0
// 	totalNumsUse := make(map[int]int)
// 	for key, v := range nums {
// 		temp += v
// 		if temp > k {
// 			break
// 		}
// 		if temp == k {
// 			newKey, err := fmt.Printf("%d%d", key, key-1)
// 			if err != nil {
// 				fmt.Println(err)
// 				panic(0)
// 			}
// 			totalNumsUse[newKey] += 1

// 			break
// 		}
// 		totalNumsUse[key] += 1
// 	}

// 	output = len(totalNumsUse)
// 	return output
// }
