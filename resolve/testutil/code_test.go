package testutil

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_isTLE(t *testing.T) {
	DebugTLE = 0
	assert.False(t, isTLE(func() {}))

	DebugTLE = time.Second
	assert.False(t, isTLE(func() {}))
	assert.True(t, isTLE(func() { select {} }))
}

type foo struct {
}

func constructor() foo {
	return foo{}
}

func (f *foo) F0(a, b int) {
	fmt.Println("f0", a, b)
}

func (f *foo) F1(a, b int) int {
	fmt.Println("f1", a, b)
	return a + b
}

func (f *foo) F2(a, b int) []int {
	fmt.Println("f2", a, b)
	return []int{a, b}
}

func TestRunClassWithFile(t *testing.T) {
	if err := RunClassWithFile(t, constructor, "code_testcase.txt", 0); err != nil {
		t.Error(err)
	}
}
