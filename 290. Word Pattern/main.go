package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abc"
	s := "b c a"
	result := wordPattern(pattern, s)
	fmt.Println("result", result)
}

func wordPattern(pattern string, s string) bool {
	s1 := strings.Split(s, " ")
	if len(pattern) != len(s1) {
		return false
	}
	p1 := []byte(pattern)
	temp1 := map[interface{}]interface{}{}
	for k, v := range p1 {
		if _, ok := temp1[v]; ok && temp1[v] != s1[k] {
			return false
		}
		if _, ok := temp1[s1[k]]; ok && temp1[s1[k]] != v {
			return false
		}
		temp1[v] = s1[k] //a:dog
		temp1[s1[k]] = v //dog:a
	}
	return true
}
