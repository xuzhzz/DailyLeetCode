package dailyleetcode

import "fmt"

func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 1000
	}
	s = s[:len(s)-2]
	s = append(s, 1111)
	s = append(s, 1111)
}

func modifySlice2(s [][]int) {
	if len(s) > 0 && len(s[0]) > 0 {
		s[0][0] = 10002
	}
}

func slice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := [][]int{s1}
	s1 = append(s1, 1)
	fmt.Printf("s1=%v, s2=%v\n", s1, s2)
	modifySlice2(s2)
	modifySlice(s1)

	fmt.Printf("s1=%v, s2=%v\n", s1, s2)
	return
}
