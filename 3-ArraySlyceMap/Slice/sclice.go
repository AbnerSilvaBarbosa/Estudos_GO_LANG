package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	s2 := a1[:3]
	fmt.Println(s1)
	fmt.Println("Slice:", s2)

}
