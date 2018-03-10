package main

import (
	"fmt"
	"sort"
)

type byLength []string

func sorting() {
	arrs := []int{34, 12, 18, 99}
	sort.Ints(arrs)
	fmt.Println("arrs:", arrs)

	str := []string{"a", "C", "z", "b"}
	sort.Strings(str)
	fmt.Println("strings:", str)

	fruits := []string{"peach", "banana", "kiwi"}

	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
