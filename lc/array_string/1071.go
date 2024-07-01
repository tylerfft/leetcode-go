package main

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}
	rst := 0
	for i := 0; i < len(str1); i++ {
		if len(str1)%(i+1) != 0 || len(str2)%(i+1) != 0 {
			continue
		}
		if ableDivided(str1, str1[:i+1]) && ableDivided(str2, str1[:i+1]) {
			rst = i + 1
		}
	}
	return str1[:rst]
}
func ableDivided(s, sub string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != sub[i%len(sub)] {
			return false
		}
	}
	return true
}
