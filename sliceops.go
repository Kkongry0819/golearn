package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func main() {
	var s []int //zero value for slice is nil
	for i := 0; i < 10; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := [...]int{1, 2}     //array
	s2 := []int{1, 2, 3}     //slice
	s3 := make([]int, 4)     //4是len
	s4 := make([]int, 4, 16) //16是cap
	fmt.Println(s1, s2, s3, cap(s3), s4, cap(s4))
	copy(s4, s2)
	fmt.Println(s4)
	s5 := []int{1, 2, 3, 4}
	s5 = append(s5, s5...)
	fmt.Println(s5)
	//front
	front := s2[0]
	s2 = s2[1:]
	//back
	back := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, back)
}
