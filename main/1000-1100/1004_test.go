package main

import (
	"testing"

	"testutil"
)

func Test_1004(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, longestOnes, "1004.txt", 0); err != nil {
		t.Fatal(err)
	}
}
