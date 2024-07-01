package main

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := hp{}
	for i := 0; i < k; i++ {
		heap.Push(&h, tuple{d: nums[i]})
	}
	for i := k; i < len(nums); i++ {
		heap.Push(&h, tuple{d: nums[i]})
		heap.Pop(&h)
	}
	return h.Top().d
}

type tuple struct{ d int }
type hp []tuple

func (h hp) Top() tuple            { return h[0] }
func (h hp) Len() int              { return len(h) }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v interface{}) { v, (*h) = (*h)[len(*h)-1], (*h)[:len(*h)-1]; return }
