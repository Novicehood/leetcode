package main

import (
	"fmt"
)

func testPanic1() {
	fmt.Println("testPanic1上半部分")
	testPanic2()
	fmt.Println("testPanic1下半部分")
}

func testPanic2() {
	fmt.Println("testPanic2上半部分")
	testPanic3()
	fmt.Println("testPanic2下半部分")
}

func testPanic3() {
	fmt.Println("testPanic3上半部分")
	panic("在testPanic3出现了panic")
	fmt.Println("testPanic3下半部分")
}

func main() {

}

func Add() (ans int) {
	a := 1
	defer func() {
		a++
	}()
	return
}

type MyStruct struct {
	MyInterface
}

// 值接收者实现 MethodA
func (s MyStruct) MethodA() {
	fmt.Println("MethodA (value receiver)")
}

// 指针接收者实现 MethodB
//func (s MyStruct) MethodB() {
//	fmt.Println("MethodB (pointer receiver)")
//}

type MyInterface interface {
	MethodA() // 由值接收者实现
	MethodB() // 由指针接收者实现
}

type MyInterface2 interface {
	MyInterface
	MethodC()
	MethodD()
}

func Func(inter MyInterface) {
	inter.MethodA()
	fmt.Println("aaaa")

}
