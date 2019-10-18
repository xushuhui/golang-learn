package main

import (
	"fmt"
	"time"
	"testing"
)
type Program interface {
	WriteHello( ) string
}
type GoProgram struct {
	
}
func (goprogram *GoProgram) WriteHello() string{
	return "fmt.Println(\"Hello\")"
}
func TestClient(t *testing.T){
	var p Program
	p = new(GoProgram)
	t.Log(p.WriteHello())
}
type IntConv  func(op int) int
func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}