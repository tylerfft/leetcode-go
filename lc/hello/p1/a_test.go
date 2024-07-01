package main

import (
	"testing"
	"testutil"
)

func Test_a(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunFuncWithFile(t, Resolve, "a.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
