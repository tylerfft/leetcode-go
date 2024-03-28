package main

func MergeAlternately(word1 string, word2 string) string {
	dd := make([]byte, len(word1)+len(word2))
	for i := 0; i < min(len(word1), len(word2)); i++ {
		dd[2*i] = word1[i]
		dd[2*i+1] = word2[i]
	}
	if len(word1) < len(word2) {
		word1, word2 = word2, word1
	}
	for i := len(word2); i < len(word1); i++ {
		dd[len(word2)+i] = word1[i]
	}
	return string(dd)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
