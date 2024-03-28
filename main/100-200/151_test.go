package main

import (
	"testing"

	"testutil/testutil"
)

func Test_1071(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, reverseWords, "151.txt", 0); err != nil {
		t.Fatal(err)
	}
}
