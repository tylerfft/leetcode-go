package main

var ss = "AEIOUaeiou"
var m = map[byte]bool{}

func init() {
	for i := 0; i < len(ss); i++ {
		m[ss[i]] = true
	}
}
func reverseVowels(s string) string {
	d := []byte(s)
	prev, back := 0, len(s)-1
	for prev < back {
		for prev < back && !m[s[prev]] {
			prev++
		}
		for prev < back && !m[s[back]] {
			back--
		}
		if prev < back {
			d[prev], d[back] = d[back], d[prev]
			prev++
			back--
		}
	}
	return string(d)
}
