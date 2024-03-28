package main

import "strings"

func reverseWords(s string) string {
	d := []string{}
	prev := 0
	back := 0
	for prev < len(s) {
		for prev < len(s) && s[prev] == ' ' {
			prev++
		}
		if prev == len(s) {
			break
		}
		back = prev + 1
		for back < len(s) && s[back] != ' ' {
			back++
		}
		d = append([]string{s[prev:back]}, d...)
		prev = back
	}
	return strings.Join(d, " ")
}
