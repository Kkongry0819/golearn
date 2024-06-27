package main

import "fmt"

func lengthOfLongestSubstring(s string) (ans int) {
	var window [128]bool // 也可以用 map，这里为了效率用的数组
	left := 0
	for right, c := range s {
		// 如果窗口内已经包含 c，那么再加入一个 c 会导致窗口内有重复元素
		// 所以要在加入 c 之前，先移出窗口内的 c
		for window[c] { // 窗口内有 c
			window[s[left]] = fa lse
			left++ // 缩小窗口
		}
		window[c] = true // 加入 c
		ans = max(ans, right-left+1) // 更新窗口长度最大值
	}
	return
}

func main() {
	m1 := map[string]string{
		"name":   "jack",
		"age":    "25",
		"course": "study",
	}
	m2 := make(map[string]string)    //empty map
	var m3 = make(map[string]string) //nil
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	for k, v := range m1 {
		fmt.Println(k, v)
		if v == "25" {
			break
		}
	}
	coursename := m1["course"]
	fmt.Println(coursename)
}
