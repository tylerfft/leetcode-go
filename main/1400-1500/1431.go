package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	rst := make([]bool, len(candies))
	maxv := 0
	for _, v := range candies {
		maxv = max(maxv, v)
	}
	for i, v := range candies {
		rst[i] = v+extraCandies >= maxv
	}
	return rst
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
