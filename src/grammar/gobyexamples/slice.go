package main

import "fmt"

func funcSlice() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	fmt.Println("append:", s)
	s = append(s, "e", "f", "g")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	sub := s[:4]
	fmt.Println("sub:", sub)
	sub = s[2:]
	fmt.Println("sub:", sub)

	t := []string{"aa", "bb", "cc"}
	fmt.Println("dcl:", t)

	s2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		s2[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			s2[i][j] = i + j
		}
	}
	fmt.Println("s2:", s2)
}
