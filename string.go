package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "kong,你玩原神吗？"
	fmt.Println(len(s))
	for _, i := range []byte(s) {
		fmt.Printf("%X\n", i)
	}
	for i, ch := range s { //ch is a rune
		fmt.Printf("%d,%X\n", i, ch)
	}
	fmt.Println("Rune count", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //解码切片中第一个utf-8编码的字符
		bytes = bytes[size:]
		fmt.Printf("%c,%d\n", ch, size)
	}
	for i, ch := range []rune(s) {
		fmt.Printf("%d,%c\n", i, ch)
	}

}
