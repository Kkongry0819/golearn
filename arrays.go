package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, arr2, arr3)
	var grid [4][3]int
	fmt.Println(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d\t", grid[i][j])
		}
		fmt.Println()
	}
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 { //i为index ,v为data
		fmt.Println(i, v)
	}
	for i, v := range grid {
		fmt.Println(i, v)
	}
	for _, v := range grid {
		fmt.Println(v)
	}
	numbers := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	maxi, maxValue := -1, -1
	for i, v := range numbers {
		if maxValue < v || maxi < -2 {
			maxi, maxValue = i, v
		}
	}
}
