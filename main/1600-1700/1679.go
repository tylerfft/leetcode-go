package main

func maxOperations(nums []int, k int) int {
	m := map[int]int{}
	ans := 0
	for _, v := range nums {
		if m[k-v] != 0 {
			m[k-v]--
			ans++
		} else {
			m[v]++
		}
	}
	return ans
}
