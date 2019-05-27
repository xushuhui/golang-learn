package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	// defer func(){
	// 	fmt.Println("final")
	// }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}

	}()
	fmt.Println("start")
	panic(errors.New("wrong"))
}
