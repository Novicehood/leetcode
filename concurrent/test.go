package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 20)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	go func() {
		for i := range ch {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i)
		}
	}()
	close(ch)
	time.Sleep(10 * time.Second)
}

func useTimer() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("finished")

}

func useWg() {
	wg1 := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		for i := 1; i < 100; i += 2 {
			wg2.Wait()
			fmt.Println(i)
			wg2.Add(1)
			wg1.Done()
		}
	}()

	go func() {
		for i := 2; i <= 100; i += 2 {
			wg1.Wait()
			fmt.Println(i)
			wg1.Add(1)
			wg2.Done()
		}
	}()
	time.Sleep(5 * time.Second)
}
