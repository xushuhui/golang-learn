package main

import (
	"testing"
)

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(a, b)
}
func TestPoint(t *testing.T) {
	a := 1
	addr := &a
	//addr = addr + 1
	t.Log(a, addr)
}
func TestString(t *testing.T) {
	var str string
	t.Log("*" + str + "*")
	t.Log(len(str))
}
