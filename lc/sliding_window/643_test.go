package main

import (
	"testing"

	"testutil"
)

func Test_643(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, findMaxAverage, "643.txt", 0); err != nil {
		t.Fatal(err)
	}
}
