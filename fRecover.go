package main

import "fmt"

func mayPanic() {
	panic("Some error")
}

func ExecRecover() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered error: ", r)
		}
	}()

	mayPanic()

	fmt.Println("After panic")
}
