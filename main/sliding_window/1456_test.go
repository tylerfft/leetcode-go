package main

import (
	"testing"

	"testutil"
)

func Test_1456(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, maxVowels, "1456.txt", 0); err != nil {
		t.Fatal(err)
	}
}
