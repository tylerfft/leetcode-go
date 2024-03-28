package main

import (
	"testing"

	"testutil"
)

func Test_1657(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, closeStrings, "1657.txt", 0); err != nil {
		t.Fatal(err)
	}
}
