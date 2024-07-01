package main

func Resolve(arr []int) []int {
	st := []int{}
	n := len(arr)
	rst := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && arr[st[len(st)-1]] <= arr[i] {
			st = st[:len(st)-1]

		}
		if len(st) != 0 {
			rst[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return rst
}
