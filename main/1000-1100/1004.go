package main

// 统计窗口内 0 的个数 cnt0，则问题转换成在 cnt0≤k 的前提下，窗口大小的最大值。

func longestOnes(nums []int, k int) int {
	ans := 0
	left := 0
	cnt0 := 0
	for right, x := range nums {
		cnt0 += 1 - x
		for cnt0 > k {
			cnt0 -= 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
