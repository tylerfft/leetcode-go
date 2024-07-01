package main

func longestSubarray(nums []int) int {
	ans := 0
	left := 0
	cnt := 0
	for right, x := range nums {
		cnt += 1 - x
		for cnt > 1 {
			cnt -= 1 - nums[left]
			left++
		}
		ans = max(ans, right-left)
	}
	return ans
}
