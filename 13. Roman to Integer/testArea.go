package main

import "strings"

type romanToIntTask struct {
	convert   map[string]int
	input     []string
	tempSlice []int
	output    int
}

func New() *romanToIntTask {
	return &romanToIntTask{
		convert: map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000},
	}
}

func romanToInt(s string) int {
	r := New()
	r.input = strings.Split(s, "")
	for _, v := range r.input {
		r.tempSlice = append(r.tempSlice, r.convert[v])
	}
	lens := len(r.tempSlice)
	r.output = r.tempSlice[lens-1]
	for i := lens - 1; i > -1; i-- {
		if i-1 < 0 {
			return r.output
		}
		if r.tempSlice[i] > r.tempSlice[i-1] {
			r.output = r.output - r.tempSlice[i-1]
		} else {
			r.output += r.tempSlice[i-1]
		}
	}
	return 0
}
