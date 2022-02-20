package main

import "strings"

func wordPattern(pattern string, s string) bool {
	strSlice := strings.Split(s, " ")
	if len([]rune(pattern)) != len(strSlice) {
		return false
	}
	wordMap := make(map[string]rune)
	patternMap := make(map[rune]string)
	for i, v := range []rune(pattern) {
		_, ok1 := patternMap[v]
		_, ok2 := wordMap[strSlice[i]]
		if !ok1 && !ok2 {
			patternMap[v] = strSlice[i]
			wordMap[strSlice[i]] = v
		}
		if patternMap[v] != strSlice[i] || wordMap[strSlice[i]] != v {
			return false
		}
	}
	return true
}
