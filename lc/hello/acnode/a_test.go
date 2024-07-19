package main

import (
	"fmt"
	"testing"
)

func Test_a(t *testing.T) {
	ac := &AcNode{}
	strs := []string{"hers", "he", "his", "she"}
	for _, v := range strs {
		ac.put(v)
	}
	ac.buildFail()

	target := "ahishers"
	o := ac
	for i, p := range target {
		o = o.son[p-'a']
		for _, l := range o.exist {
			fmt.Println(i, target[i-l+1:i+1])
		}
	}
}
