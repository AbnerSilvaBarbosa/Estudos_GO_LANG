package main

import "fmt"

func main() {

	a1 := [10]int{1, 2, 3, 4, 23, 17, 18, 20, 9}

	s := make([]int, 10, 20)
	s = append(s, a1[2])

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(s, len(s), cap(s))

}
