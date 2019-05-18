package test

import (
	"testing"
)

const (
	Mon = iota + 1
	Tues
	Wed
)
const (
	Read = 1 << iota
	Write
	Execut
)

func TestConstant(t *testing.T) {
	a := 7
	t.Log(Mon, Tues)
	t.Log(a&Read == Read, a&Write == Write, a&Execut == Execut)
}
