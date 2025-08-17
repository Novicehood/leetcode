package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	Val byte
}

func main() {
	var s S
	fmt.Println(unsafe.Sizeof(s))
}
