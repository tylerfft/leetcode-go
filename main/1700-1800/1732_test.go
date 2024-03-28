package main

import (
	"testing"

	"testutil"
)

func Test_1732(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, largestAltitude, "1732.txt", 0); err != nil {
		t.Fatal(err)
	}
}
