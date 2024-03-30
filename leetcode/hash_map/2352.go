package main

func equalPairs(grid [][]int) int {
	mr := map[[200]int]int{}
	r, c := len(grid), len(grid[0])
	for rr := 0; rr < r; rr++ {
		str := [200]int{}
		for cc := 0; cc < c; cc++ {
			str[cc] = grid[rr][cc]
		}
		mr[str]++
	}
	ans := 0
	for cc := 0; cc < c; cc++ {
		str := [200]int{}
		for rr := 0; rr < r; rr++ {
			str[rr] = grid[rr][cc]
		}
		ans += mr[str]
	}
	return ans
}
