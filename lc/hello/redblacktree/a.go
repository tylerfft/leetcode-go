package main

import (
	"container/heap"
	"context"
	"fmt"
	"time"
)

func Resolve(str string) string {
	return ""
}

// 1 node is red or balck
// 2 root is black
// 3 leaf is black
// 4 one node is red, its two children are black
// 5 black node count from root to leaf is same

func users(ctx context.Context) {
	heap.Push()
	// 设置请求的截止时间为当前时间加上 1 秒钟
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	// 启动一个 goroutine 来处理请求
	go func(ctx context.Context) {
		// 等待请求完成或者超时
		select {
		case <-time.After(time.Millisecond * 500):
			// 请求完成
			fmt.Println("Request finish")
		case <-ctx.Done():
			// 请求超时或者被取消
			fmt.Println("Request canceled or timed out")
		}
	}(ctx)

	// 等待一段时间后取消请求
	time.Sleep(time.Millisecond * 1500)
	cancel()
}
