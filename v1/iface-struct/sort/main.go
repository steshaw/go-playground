package main

import (
	"fmt"
	"sort"
)

type intList struct {
	array []int
}

func (l intList) Len() int {
	result := len(l.array)
	fmt.Printf("Len => %v\n", result)
	return result
}

func (l intList) Less(i, j int) bool {
	result := l.array[i] < l.array[j]
	fmt.Printf("Less %d %d => %v\n", i, j, result)
	return result
}

func (l intList) Swap(i, j int) {
	li := l.array[i]
	lj := l.array[j]
	l.array[i], l.array[j] = l.array[j], l.array[i]
	fmt.Printf(
		"Swap %d:%d %d:%d => %v\n",
		i, li, j, lj, l.array,
	)
}

func main() {
	l := intList{array: []int{4, 5, 2, 8, 1, 9, 3}}
	fmt.Println(l)
	sort.Sort(l)
	fmt.Println(l)
}
