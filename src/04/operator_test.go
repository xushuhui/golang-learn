package main

import (
	"testing"
)

const (
	Read = 1 << iota
	Write
	Execut
)

func TestConstant(t *testing.T) {
	a := 7
	a = a &^ Read
	t.Log(a&Read == Read, a&Write == Write, a&Execut == Execut)
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 5, 4}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == d)
}
