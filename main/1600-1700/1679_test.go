package main

import (
	"testing"

	"testutil"
)

func Test_1679(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, maxOperations, "1679.txt", 0); err != nil {
		t.Fatal(err)
	}
}
