package main

import "os"

func ExecPanic() {
	panic("Some error")

	_, err := os.Create("test/panicfile")
	if err != nil {
		panic(err)
	}
}
