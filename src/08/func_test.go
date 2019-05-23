package main

import (
	"fmt"
	"math/rand"
	"time"

	"testing"
)

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
func slowFunc(op int) int{
	time.Sleep(time.Second*1)
	return op
}
func sum(ops ...int)int{
	ret :=0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestFn(t *testing.T) {
	a, _ := returnMultiValue()
	t.Log(a)
	tsfunc := timeSpent(slowFunc)
	t.Log(tsfunc(1))
	t.Log(sum(1,2,3))
}
func clear(){
	fmt.Println("clear")
}
func TestDefer(t *testing.T){
	defer clear()
	fmt.Println("start")
	panic("error")
}