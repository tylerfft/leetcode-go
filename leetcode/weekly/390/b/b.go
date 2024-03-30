package main

import "math"

// from miao
// func minOperations(k int) int {
// 	ans := math.MaxInt
// 	for m := 1; m <= k; m++ {
// 		ans = min(ans, m-1+(k-1)/m)
// 	}
// 	return ans
// }

func minOperations(k int) int {
	ans := math.MaxInt
	for m := 1; m <= k; m++ {
		ans = min(ans, m-1+(k-1)/m)
	}
	return ans
}
