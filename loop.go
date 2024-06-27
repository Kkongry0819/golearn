package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	if n > 0 {
		for ; n > 0; n /= 2 {
			lsb := n % 2
			result = strconv.Itoa(lsb) + result
		}
		return result
	} else if n == 0 {
		return "0"
	} else {
		//panic(err)
		return "error"
	}
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func forever() {
	for {
		fmt.Println("loop")
	}
}
func choose(n string) {
	if n == "/" {
		q, r := div(2, 3)
		fmt.Println(q, r)
	}
}
func div(n, m int) (q, r int) {
	q = n / m
	r = n % m
	return
}
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args %d,%d\n", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
func swap(a *int, b *int) {
	*a, *b = *b, *a
}
func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1011-->1101
		convertToBin(0),
		convertToBin(-1),
	)
	printFile("test.txt")
	//forever()
	choose("/")
	fmt.Printf("result is %d\n", apply(pow, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
