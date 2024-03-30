package main

func maxArea(height []int) int {
	prev, back := 0, len(height)-1
	minv := min(height[0], height[len(height)-1])
	ans := minv * (back - prev)
	for prev < back {
		for prev < back && height[prev] <= minv {
			prev++
		}
		for prev < back && height[back] <= minv {
			back--
		}
		if prev < back {
			minv = min(height[prev], height[back])
			ans = max(ans, (back-prev)*minv)
		}
	}
	return ans
}
