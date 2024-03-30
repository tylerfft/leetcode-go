package main

// 从左到右遍历数组，能种花就立刻种花。
// 如何判断能否种花？由于「花不能种植在相邻的地块上」，如果要在下标 iii 处种花，需要满足 flowerbed[i−1],flowerbed[i],flowerbed[i+1]\textit{flowerbed}[i-1],\textit{flowerbed}[i],\textit{flowerbed}[i+1]flowerbed[i−1],flowerbed[i],flowerbed[i+1] 均为 000。

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append(append([]int{0}, flowerbed...), 0)
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
