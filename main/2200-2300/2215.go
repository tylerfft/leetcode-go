package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := map[int]bool{}
	m2 := map[int]bool{}
	for _, x := range nums1 {
		m1[x] = true
	}
	for _, x := range nums2 {
		m2[x] = true
	}
	ans := make([][]int, 2)
	for x, _ := range m1 {
		if !m2[x] {
			ans[0] = append(ans[0], x)
		}
	}
	for x, _ := range m2 {
		if !m1[x] {
			ans[1] = append(ans[1], x)
		}
	}
	return ans
}
