package main

import (
	"testing"

	"testutil/testutil"
)

func Test_1768(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, MergeAlternately, "1768.txt", 0); err != nil {
		t.Fatal(err)
	}
}
