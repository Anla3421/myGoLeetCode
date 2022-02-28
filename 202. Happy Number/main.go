package main

import "fmt"

func main() {
	nums := 19
	result := isHappy(nums)
	fmt.Println("result", result)
}

func isHappy(n int) bool {
	for i := 51; i <= 100; i++ {
		n = i
		input := n * n
		temp := numSplit(input)
		for {
			v1 := 0
			for _, v := range temp {
				v1 += v * v
			}
			// fmt.Println(v1)
			if v1 <= 10 {
				fmt.Printf("input:%v,%v\n", n, v1)
				// fmt.Println(v1)
				// fmt.Println("this is end")
				// return true
				break
			}
			temp = numSplit(v1)
		}
	}
	return true
}

func numSplit(input int) (result []int) {
	i := 0
	for {
		result = append(result, input%10)
		i++
		if input/10 == 0 {
			break
		}
		input = input / 10
	}
	return result
}
