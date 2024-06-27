package main

import "fmt"

func updateSlice(arr []int) {
	arr[0] = 100
}
func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	s := arr[2:4]
	fmt.Println(s)
	fmt.Println(arr[:4]) //切片是视图
	fmt.Println(arr[:])
	updateSlice(arr[:])
	fmt.Println(arr)
	s = arr[:2] //reslice
	s1 := arr[3:4]
	s2 := s1[0:2]   //通过s1的相对坐标再对arr的切片 ptr len cap 指针 长度 向后的容量
	fmt.Println(s1) //slice可以向后扩展
	fmt.Printf("s1 type: %T\n", s1)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s1: %d\n", len(s1))
	fmt.Printf("s1: %d\n", cap(s1))
	fmt.Println(s2)
	fmt.Println(s) //可以对切片再切片
	s3 := append(s1, 6)
	fmt.Println(s3)
	s4 := append(s3, 7)
	fmt.Println(s4)
	fmt.Println(arr) //当slice的长度大于cap，slice视图参照一个cap更大的数组
}
