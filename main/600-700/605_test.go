package main

import (
	"testing"

	"testutil/testutil"
)

func Test_605(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, canPlaceFlowers, "605.txt", 0); err != nil {
		t.Fatal(err)
	}
}
