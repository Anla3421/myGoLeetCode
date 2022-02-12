package main

// func addDigits(num int) int {
// 	if num < 10 {
// 		return num
// 	}

// 	if r := num % 9; r == 0 {
// 		return 9
// 	} else {
// 		return r
// 	}
// }

// func addDigits(num int) int {
// 	for num >= 10 {
// 		tmpN := 0
// 		for int(math.Pow(10, float64(tmpN))) <= num {
// 			tmpN += 1
// 			fmt.Println(tmpN)
// 		}
// 		tmpN -= 1
// 		slice := []int{}
// 		for i := tmpN; i >= 0; i-- {
// 			divisor := int(math.Pow(10, float64(i)))
// 			splitOne := (num) / divisor
// 			num = (num) % divisor

// 			slice = append(slice, splitOne)
// 		}
// 		num = 0
// 		for _, eachOne := range slice {
// 			num += eachOne

// 		}
// 		fmt.Println(num)
// 	}
// 	return num
// }

// func addDigits(num int) int {
// 	for num > 9 {
// 		tmp_slice := []int{}

// 		for num > 9 {
// 			fmt.Println(num)
// 			tmp_slice = append(tmp_slice, num%10)
// 			num = num / 10
// 			if num <= 9 {
// 				tmp_slice = append(tmp_slice, num)
// 			}
// 		}

// 		num = 0
// 		fmt.Println(tmp_slice)
// 		for _, each := range tmp_slice {
// 			num += each
// 		}

// 	}
// 	return num
// }

// func addDigits(num int) int {
// 	for num > 9 {
// 		tmp_slice := []int{}

// 		for num > 9 {
// 			tmp_slice = append(tmp_slice, num%10)
// 			num = num / 10
// 			if num <= 9 {
// 				tmp_slice = append(tmp_slice, num)
// 			}
// 		}

// 		num = 0
// 		for _, each := range tmp_slice {
// 			num += each
// 		}
// 	}
// 	return num
// }
