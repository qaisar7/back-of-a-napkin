/*
  Back of a napkin test of how channels communicate with each other.
  Mocks testing a request being processed by multiple channels and preemptively abort when successful.
  Closes all the threads before exiting.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	c := make(chan int, 1)
	threads := []chan int{}
	shutdown := make(chan struct{})
	for i := 0; i < 6; i++ {
		t := make(chan int)
		threads = append(threads, t)
		go func(i int, c chan int, t chan int) {
			fmt.Printf("Thread %d sleeping now\n", i)
			tm := time.Now()
			d := time.Duration(rand.Int63n(6000))
			fmt.Println("duration ", time.Millisecond*d)
			select {
			case <-time.After(time.Millisecond * d):
				fmt.Printf("Thread %d waking after %s \n", i, time.Now().Sub(tm))
				c <- i
			case <-shutdown:
				fmt.Printf("Thread %d shutting down\n", i)
				close(c)
			}

			t <- i
			close(t)
		}(i, c, t)
	}

	go func() {
		for _, t := range threads {
			for i := range t {
				fmt.Printf("thread %d closed\n", i)
			}
		}
		close(c)
	}()

	for i := range c {
		fmt.Printf("Thread %d done\n", i)
		if i == 3 {
			close(shutdown)
			fmt.Println(" got result, closing")
			break
		}
	}
}
