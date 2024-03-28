package main

import (
	"testing"

	"testutil"
)

func Test_724(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, pivotIndex, "724.txt", 0); err != nil {
		t.Fatal(err)
	}
}
