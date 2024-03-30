package main

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	is := 0
	for it := 0; it < len(t); it++ {
		if s[is] == t[it] {
			is++
			if is == len(s) {
				return true
			}
		}
	}
	return false
}
