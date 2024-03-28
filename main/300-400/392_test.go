package main

import (
	"testing"

	"testutil"
)

func Test_392(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, isSubsequence, "392.txt", 0); err != nil {
		t.Fatal(err)
	}
}
