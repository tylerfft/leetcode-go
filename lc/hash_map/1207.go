package main

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, v := range arr {
		m[v]++
	}
	mm := map[int]bool{}
	for _, v := range m {
		if mm[v] {
			return false
		}
		mm[v] = true
	}
	return true
}
