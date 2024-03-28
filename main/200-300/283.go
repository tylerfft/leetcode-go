package main

func moveZeroes(nums []int) {
	prev := 0
	for back := 0; back < len(nums); back++ {
		if nums[back] != 0 {
			nums[prev] = nums[back]
			if prev != back {
				nums[back] = 0
			}
			prev++
		}
	}
}
