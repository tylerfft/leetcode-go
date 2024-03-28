package common

const INF = 1e15

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
