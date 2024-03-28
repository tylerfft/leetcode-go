package main

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	obj := 0
	for i, v := range nums {
		if obj == sum-obj-v {
			return i
		}
		obj += v
	}
	return -1
}
