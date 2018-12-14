/*
  Back of a napkin test of how channels communicate with each other.
  Mocks testing a request being processed by multiple channels.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generateInput generates a slice of size 'n' each having a random time duration number between 0-6000.
func generateInput(n int) []time.Duration {
	ar := []time.Duration{}
	for i := 0; i < n; i++ {
		d := time.Duration(rand.Int63n(6000))
		ar = append(ar, d)
	}
	return ar
}

type threadId struct {
	id   int
	jobs int
}

func main() {
	startTime := time.Now()
	waits := generateInput(10)
	fmt.Println("channnnnnnels", waits)
	totalThreads := 3
	jobC := make(chan time.Duration, totalThreads)
	//shutdown := make(chan struct{})
	go func() {
		for _, w := range waits {
			fmt.Printf("sending wait %s on job channel\n", time.Millisecond*w)
			jobC <- w
		}
		fmt.Println("closing job channel")
		close(jobC)
	}()

	threads := []chan threadId{}

	for i := 0; i < totalThreads; i++ {
		t := make(chan threadId)
		threads = append(threads, t)
		go func(i int, c chan time.Duration, t chan threadId) {
			jobs := 0
			for d := range c {
				fmt.Printf("Thread %d sleeping now for %s \n", i, time.Millisecond*d)
				tm := time.Now()
				time.Sleep(time.Millisecond * d)
				fmt.Printf("Thread %d waking after %s \n", i, time.Now().Sub(tm))
				jobs++
			}
			t <- threadId{i, jobs}
			close(t)
		}(i, jobC, t)
	}

	for _, t := range threads {
		for th := range t {
			fmt.Printf("thread %d closed with jobs %d\n", th.id, th.jobs)
		}
	}
	var totalWaits time.Duration
	for _, w := range waits {
		totalWaits += w
	}
	fmt.Printf("total program time %s for total waits %s\n", time.Now().Sub(startTime), time.Millisecond*totalWaits)
}
