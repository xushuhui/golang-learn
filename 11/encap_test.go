package main

import (
	"unsafe"
	"fmt"
	"testing"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

func TestCreateObj(t *testing.T) {
	e := Employee{"1", "X", 18}
	e1 := Employee{Name: "Y", Age: 20}
	e2 := new(Employee)
	e2.ID = "2"
	e2.Name = "Z"
	e2.Age = 25
	t.Log(e)
	t.Log(e1)
	t.Log(e1.ID)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)

}
// func (e Employee) String() string{
//	fmt.Printf("addresss %x",unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.ID,e.Name,e.Age)
// }
func (e *Employee) String() string{
	fmt.Printf("addresss %x",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.ID,e.Name,e.Age)
}
func TestStruct(t *testing.T){
	e := Employee{"1", "X", 18}
	fmt.Printf("addresss %x",unsafe.Pointer(&e.Name))
	t.Log(e.String())
}