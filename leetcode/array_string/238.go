package main

func productExceptSelf(nums []int) []int {
	mul := nums[0]
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = mul
		mul *= nums[i]
	}
	mul = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		ans[i] *= mul
		mul *= nums[i]
	}
	return ans
}
