package main

import (
	"fmt"
	"os"
)

func grade(score int) string { //不break用fallthrough
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
		//default:
		//	panic(fmt.Sprintf("wrong score: %d", score))
	}
	return g
}
func main() {
	const filename = "test.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
	fmt.Println(
		grade(0),
		grade(60),
		grade(20),
		grade(24),
		grade(32),
		grade(99),
		grade(100))
}
