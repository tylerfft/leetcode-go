package main

import (
	"testing"

	"testutil"
)

func Test_215(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, findKthLargest, "215.txt", 0); err != nil {
		t.Fatal(err)
	}
}
