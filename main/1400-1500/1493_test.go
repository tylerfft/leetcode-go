package main

import (
	"testing"

	"testutil"
)

func Test_1493(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, longestSubarray, "1493.txt", 0); err != nil {
		t.Fatal(err)
	}
}
