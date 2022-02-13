package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "MCMXCIV"
	// s := "CIV"
	// s := "IXXX"
	result := romanToInt(s)
	fmt.Println("result", result)
}

func romanToInt(s string) int {
	var (
		input     []string
		tempSlice []int
		output    int
	)
	convert := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	input = strings.Split(s, "")
	for _, v := range input {
		tempSlice = append(tempSlice, convert[v])
	}
	lens := len(tempSlice)
	output = tempSlice[lens-1]
	for i := lens - 1; i > -1; i-- {
		if i-1 < 0 {
			return output
		}
		if tempSlice[i] > tempSlice[i-1] {
			output = output - tempSlice[i-1]
		} else {
			output += tempSlice[i-1]
		}
	}
	return 0
}
