package main

import (
	"testing"

	"testutil"
)

func Test_283(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, moveZeroes, "283.txt", 0); err != nil {
		t.Fatal(err)
	}
}
