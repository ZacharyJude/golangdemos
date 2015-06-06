package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	scale := 100000
	c := make(chan int, 32)

	for i := 0; i < scale; i++ {

		wg.Add(1)

		go func(c chan int) {
			for i := 0; i < 10; i++ {
				c <- i * i
			}
			wg.Done()
		}(c)

		wg.Add(1)

		go func(c chan int) {
			for i := 0; i < 10; i++ {
				s := <-c
				fmt.Println(s)
			}
			wg.Done()
		}(c)
	}

	wg.Wait()
}
