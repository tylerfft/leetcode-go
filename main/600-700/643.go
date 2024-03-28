package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	ans := -math.MaxFloat64
	sum := 0.0
	for i := 0; i < k-1; i++ {
		sum += float64(nums[i])
	}
	for i := k - 1; i < len(nums); i++ {
		sum += float64(nums[i])
		ans = max(ans, sum/float64(k))
		sum -= float64(nums[i-k+1])
	}
	return ans
}
