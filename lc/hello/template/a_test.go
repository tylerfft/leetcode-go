package main

import (
	"testing"
	"testutil"
)

func Test_a(t *testing.T) {
	if err := testutil.RunFuncWithFile(t, Resolve, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}
