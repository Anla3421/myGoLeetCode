package main

func twoSum(nums []int, target int) []int {
	s := newSet()

	for idx, num := range nums {
		desired := target - num
		if s.has(desired) {
			idx2 := s.get(desired)
			return []int{idx2, idx}
		}
		s.add(num, idx)
	}
	return []int{}
}

type set struct {
	items map[int]int
}

func newSet() *set {
	return &set{
		items: make(map[int]int),
	}
}

func (s *set) add(value, position int) {
	s.items[value] = position
}

func (s *set) has(i int) bool {
	_, exists := s.items[i]
	return exists
}

func (s *set) get(i int) int {
	idx, exists := s.items[i]
	if !exists {
		return -1
	}
	return idx
}

// func twoSum(nums []int, target int) []int {
// 	lens := len(nums)
// 	for i := 0; i < lens; i++ {
// 		for j := i + 1; j < lens; j++ {
// 			if target-nums[i] == nums[j] {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// func twoSum(nums []int, target int) []int {
// 	mapNums := map[int]int{}
// 	var result []int
// 	for index, value := range nums {
// 		mapNums[value] = index
// 	}
// 	for k, v := range mapNums {
// 		if _, ok := mapNums[target-k]; ok {
// 			if mapNums[target-k] != v {
// 				return []int{v, mapNums[target-k]}
// 			}
// 			for index, value := range nums {
// 				if value == target-k {
// 					result = append(result, index)
// 				}
// 				if len(result) == 2 {
// 					return result
// 				}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// func twoSum(nums []int, target int) []int {
// 	prefix := map[int]int{0: 1}
// 	var sum int
// 	var count int
// 	var result []int
// 	for k, v := range nums {
// 		sum += v
// 		count += prefix[sum-target]
// 		if prefix[sum-target] != 0 {
// 			result = []int{k - 1, k}
// 		}
// 		prefix[sum]++
// 	}
// 	// fmt.Println(prefix)
// 	// fmt.Println(count)
// 	// fmt.Println(result)
// 	return result
// }
