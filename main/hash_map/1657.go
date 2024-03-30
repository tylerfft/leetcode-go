package main

func closeStrings(word1 string, word2 string) bool {
	m1 := map[byte]int{}
	m2 := map[byte]int{}
	for i := 0; i < len(word1); i++ {
		m1[word1[i]]++
	}
	for i := 0; i < len(word2); i++ {
		m2[word2[i]]++
	}
	cnt := map[int]int{}
	for k, v := range m1 {
		if m2[k] == 0 {
			return false
		}
		cnt[v]++
	}
	for k, v := range m2 {
		if m1[k] == 0 {
			return false
		}
		cnt[v]--
	}
	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
