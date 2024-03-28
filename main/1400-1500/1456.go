package main

var str = "AEIOUaeiou"
var set = map[byte]bool{}

func init() {
	for i := 0; i < len(str); i++ {
		set[str[i]] = true
	}
}

func maxVowels(s string, k int) int {
	ans := 0
	cnt := 0
	for i := 0; i < k-1; i++ {
		if set[s[i]] {
			cnt++
		}
	}
	for i := k - 1; i < len(s); i++ {
		if set[s[i]] {
			cnt++
		}
		ans = max(ans, cnt)
		if set[s[i-k+1]] {
			cnt--
		}
	}
	return ans
}
