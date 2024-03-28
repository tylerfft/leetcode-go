package main

import (
	"testing"

	"testutil"
)

func Test_1207(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, uniqueOccurrences, "1207.txt", 0); err != nil {
		t.Fatal(err)
	}
}
