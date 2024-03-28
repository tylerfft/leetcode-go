package main

import (
	"testing"

	"testutil"
)

func Test_2352(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, equalPairs, "2352.txt", 0); err != nil {
		t.Fatal(err)
	}
}
