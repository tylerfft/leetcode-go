package main

func largestAltitude(gain []int) int {
	ans := 0
	pre := 0
	for i := 0; i < len(gain); i++ {
		pre += gain[i]
		ans = max(ans, pre)
	}
	return ans
}
