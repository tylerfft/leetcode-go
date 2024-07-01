package main

import (
	"testing"

	"testutil"
)

func Test_443(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, compress, "443.txt", 0); err != nil {
		t.Fatal(err)
	}
}
