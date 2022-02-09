package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := 199
	result := addDigits(num)
	fmt.Println("result:", result)
}

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	var total int
	temp := fmt.Sprint(num)
	for {
		s := strings.Split(temp, "")
		total = 0
		for _, v := range s {
			i, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				panic(0)
			}
			total = total + i
		}
		newTotal := (total/10 + total%10)
		if newTotal < 10 {
			return newTotal
		}
		temp = strconv.Itoa(newTotal)
	}
}
