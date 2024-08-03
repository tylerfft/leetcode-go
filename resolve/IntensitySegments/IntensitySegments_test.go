package main

import (
	"testing"
	"testutil"
)

func Test_a(t *testing.T) {
	if err := testutil.RunClassWithFile(t, NewIntensitySegments, "testcase.txt", 0); err != nil {
		t.Fatal(err)
	}
}
