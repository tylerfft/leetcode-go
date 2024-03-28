package main

import (
	"testing"

	"testutil"
)

func Test_2215(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, findDifference, "2215.txt", 0); err != nil {
		t.Fatal(err)
	}
}
