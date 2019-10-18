package main

import (
	"errors"
	"testing"
)

var LessError error = errors.New("n Less")
var LargeError error = errors.New("n large")

func GetFib(n int) ([]int, error) {
	if n < 2 {
		return nil, LessError
	}
	if n > 100 {
		return nil, LargeError
	}
	if n < 2 || n > 10 {
		return nil, errors.New("n be in [2,100]")
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
func TestGetFib(t *testing.T) {
	// if v, err := GetFib(-10); err != nil {
	// 	t.Error(err)
	// } else {
	// 	t.Log(v)
	// }
	v, err := GetFib(-10)
	if err == LessError {
		t.Error("need large")
	}
	if err == LargeError {
		t.Error("need less")
	}
	t.Log(v)
}
