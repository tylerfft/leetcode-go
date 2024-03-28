package main

import (
	"testing"

	"testutil/testutil"
)

func Test_238(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, productExceptSelf, "238.txt", 0); err != nil {
		t.Fatal(err)
	}
}
