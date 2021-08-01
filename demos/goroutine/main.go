package main

import (
	"fmt"
	"sync"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		compute(4)
		wg.Done()
	}()
	go func() {
		compute(6)
		wg.Done()
	}()

	wg.Wait()
}
