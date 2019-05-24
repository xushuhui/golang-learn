package main

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Log(len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Log(len(m2))
	m3 := make(map[int]int, 10)
	t.Log(len(m3))
}
func TestAccessNotExistKey(t *testing.T) {
	m2 := map[int]int{}
	t.Log(m2[1])
	m2[2] = 0
	t.Log(m2[2])
	m2[3] = 0
	if val, ok := m2[3]; ok {
		t.Log("exist", val)
	} else {
		t.Log("no exist")
	}
}
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
func TestMapFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}
func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	set[1] = true
	n := 1
	if set[n] {
		t.Log("exist")
	} else {
		t.Log("no exist")
	}
	set[3] = true
	t.Log(len(set))
	delete(set, 2)
	t.Log(set)
}
