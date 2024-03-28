package main

import (
	"testing"

	"testutil/testutil"
)

func Test_334(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, increasingTriplet, "334.txt", 0); err != nil {
		t.Fatal(err)
	}
}
