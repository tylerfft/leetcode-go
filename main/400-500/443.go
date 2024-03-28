package main

import "fmt"

func compress(chars []byte) int {
	ans := 0
	for prev := 0; prev < len(chars); {
		cur := prev
		for cur < len(chars) && chars[cur] == chars[prev] {
			cur++
		}
		chars[ans] = chars[prev]
		ans++
		if cur > prev+1 {
			cnt := fmt.Sprint(cur - prev)
			for j := 0; j < len(cnt); j++ {
				chars[ans] = cnt[j]
				ans++
			}
		}
		prev = cur
	}
	return ans
}
