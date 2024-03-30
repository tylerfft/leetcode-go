package main

import (
	"testing"

	"testutil"
)

func Test_1431(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, kidsWithCandies, "1431.txt", 0); err != nil {
		t.Fatal(err)
	}
}
