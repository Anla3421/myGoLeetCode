package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "cacc"
	t := "aacc"
	result := isAnagram(s, t)
	fmt.Println("result", result)
}

func isAnagram(s string, t string) bool {
	temp := map[string]int{}
	s1 := strings.Split(s, "")
	t1 := strings.Split(t, "")
	if len(s1) != len(t1) {
		return false
	}
	for k, v := range s1 {
		temp[v]++
		temp[t1[k]]--
	}
	for _, v := range temp {
		if v != 0 {
			return false
		}
	}
	return true
}
