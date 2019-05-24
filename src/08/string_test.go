package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T)  {
	str := "A,B,C"
	slice := strings.Split(str,",")
	for _, v := range slice {
		t.Log(v)
	}
	t.Log(strings.Join(slice,"-"))
}
func TestConv(t *testing.T){
	s:=strconv.Itoa(10)
	t.Log(s)
	if i,err := strconv.Atoi("10"); err == nil{
		t.Log(10+i)
	}
	
}