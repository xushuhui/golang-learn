package main

import (
	"fmt"
	"testing"
)

type Code string
type Program interface {
	WriteHello() Code
}
type GoProgram struct {
}
type JavaProgram struct {
}

func (p *GoProgram) WriteHello() Code {
	return "fmt.Println(\"Hello\")"
}
func (p *JavaProgram) WriteHello() Code {
	return "System.our.Println(\"Hello\")"
}
func writeFirst(p Program) {
	fmt.Printf("%T %v\n", p, p.WriteHello())
}
func TestWrite(t *testing.T) {
	goPro := new(GoProgram)
	javPro := new(JavaProgram)
	writeFirst(goPro)
	writeFirst(javPro)
}
