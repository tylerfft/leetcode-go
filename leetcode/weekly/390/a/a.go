package main

// from miao
func maximumLengthSubstring(s string) (ans int) {
	cnts := [26]int{}
	left := 0
	for right, x := range s {
		cnts[x-'a']++
		for cnts[x-'a'] > 2 {
			cnts[s[left]-'a']--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
