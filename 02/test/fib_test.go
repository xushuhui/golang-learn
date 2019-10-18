package test

import (
	"testing"
)

func TestFib(t *testing.T) {
	a := 1
	b := 2
	c, d := 3, 4
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log(c, d)
	t.Log("Test")
}
