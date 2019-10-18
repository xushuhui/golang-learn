package main

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}
func TestArrayTravel(t *testing.T) {
	arr1 := [4]int{1, 2, 3, 4}
	for i := 0; i < len(arr1); i++ {
		t.Log(arr1[i])
	}
	for k, v := range arr1 {
		t.Log(k, v)
	}
	for _, v := range arr1 {
		t.Log(v)
	}

}
func TestArraySection(t *testing.T) {
	arr1 := [4]int{1, 2, 3, 4}
	t.Log(arr1[:3])
}
func TestSliceInit(t *testing.T) {
	var slice []int
	t.Log(len(slice), cap(slice))
	slice = append(slice, 1)
	t.Log(len(slice), cap(slice))

	slice1 := []int{1, 2, 3, 4}
	t.Log(len(slice1), cap(slice1))

	slice2 := make([]int, 3, 5)
	t.Log(len(slice2), cap(slice2))
	t.Log(slice2[0], slice2[1], slice2[2])
	slice2 = append(slice2, 1)
	t.Log(slice2[0], slice2[1], slice2[2], slice2[3])
	t.Log(len(slice2), cap(slice2))
}
func TestSliceGrow(t *testing.T) {
	s := []int{}
	for i := 0; i < 0; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
func TestSliceShareMemory(t *testing.T) {
	year := []string{"J", "F", "M", "A", "Ma", "JUN", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	y2 := year[3:6]
	t.Log(y2, len(y2), cap(y2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknow"
	t.Log(year)
}
func TestSliceCompare(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{1, 2, 3, 4}
	if slice1 == slice2 {
		t.Log("yes")
	}
}
