package main

import (
	"testing"

	"testutil"
)

func Test_1071(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, gcdOfStrings, "1071.txt", 0); err != nil {
		t.Fatal(err)
	}
}
