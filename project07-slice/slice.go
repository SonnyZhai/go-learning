package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a) // [     0 1 2 3 4 5 6 7 8 9]

	// var b = [...]int{3, 7, 8, 9, 1} 进行排序，sort包
	var b = [...]int{3, 7, 8, 9, 1}
	sort.Ints(b[:])
	fmt.Println(b) // [1 3 7 8 9]

	s := make([]string, 3)
	fmt.Println("emp:", s) // emp: [  ]

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)    // set: [a b c]
	fmt.Println("get:", s[2]) // get: c

	fmt.Println("len:", len(s)) // len: 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) // apd: [a b c d e f]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) // cpy: [a b c d e f]

	l := s[2:5]
	fmt.Println("sl1:", l) // sl1: [c d e]

	l = s[:5]
	fmt.Println("sl2:", l) // sl2: [a b c d e]

	l = s[2:]
	fmt.Println("sl3:", l) // sl3: [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // dcl: [g h i]

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0] [1 2] [2 3 4]]

	// 通过 buffer 串联字符串
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok {
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Println(buffer.String())

}

func getNextString() (string, bool) {
	return "", false
}
