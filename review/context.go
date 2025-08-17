package main

import (
	"context"
	"fmt"
	"time"
)

type Ctx interface {
	Read(string)
	Print(string)
}

type P struct {
	Ctx
	name string
	age  int
}

func (p P) Read(msg string) {
	fmt.Println(msg)
}

func main() {
	cancelCtx, _ := context.WithCancel(context.Background())
	timeCtx, _ := context.WithTimeout(cancelCtx, time.Second*3)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("subroutine exiting...")
				return
			default:
				fmt.Println("doing...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(timeCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("done")
}
