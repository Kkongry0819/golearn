package main

import (
	"fmt"
	"test1/queue"
)

func main() {
	q := new(queue.Queue)
	var q1 = queue.Queue{1, 2, 3}
	q1.Push(6)
	fmt.Println(q1)
	q.Push(12)
	q.Push(13)
	q.Push(14)
	q.Push(15)
	fmt.Println(q.Pop())
	for _, d := range *q {
		fmt.Printf("%d ", d)
	}
	fmt.Println()
	fmt.Println(q.Len())
}
