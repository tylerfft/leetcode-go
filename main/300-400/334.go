package main

import (
	. "common"
)

func increasingTriplet(nums []int) bool {
	dd := []int{INF, INF, INF}
	for _, v := range nums {
		for i := 0; i < 3; i++ {
			if dd[i] >= v {
				dd[i] = v
				break
			}
		}
	}
	return dd[0] != INF && dd[1] != INF && dd[2] != INF
}
