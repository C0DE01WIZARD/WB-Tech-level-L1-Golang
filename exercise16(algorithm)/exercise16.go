package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Exercise 16\n")
	slc := []int{10, 9, 5, 6, 8, 7, 3, 4, 2, 1}
	sort.Slice(slc, func(i, j int) bool { return slc[j] > slc[i] })
	fmt.Printf("Сортировка слайса slc: %v\n", slc)
}