package main

import (
	"sort"
)

func Resolve(arr [][]int) []int {
	m, n := len(arr), len(arr[0])
	bfs := func(i, j int) int {
		cnt := 0
		q := [][]int{{i, j}}
		for len(q) > 0 {
			node := q[0]
			cnt++
			q = q[1:]
			for _, dir := range dirs {
				ii, jj := node[0]+dir[0], node[1]+dir[1]
				if ii < 0 || ii >= m || jj < 0 || jj >= n {
					continue
				}
				if arr[ii][jj] == 0 {
					q = append(q, []int{ii, jj})
					arr[ii][jj] = -1
				}
			}
		}
		return cnt
	}
	set := map[int]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				arr[i][j] = -1
				cnt := bfs(i, j)
				set[cnt] = true
			}
		}
	}
	rst := []int{}
	for key := range set {
		rst = append(rst, key)
	}
	sort.Ints(rst)
	return rst
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// [[0,0,0],[0,0,0],[0,0,0]]
// [9]

// [[0,0,0],[0,0,0],[0,0,0]]
// [8]

// [[0,1,0],
// [1,1,1],
// [0,1,0]]
// [1,1,1,1]
