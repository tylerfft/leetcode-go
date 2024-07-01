package main

import (
	"testing"

	"testutil"
)

func Test_151(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, maxArea, "11.txt", 0); err != nil {
		t.Fatal(err)
	}
}
