package main

import (
	"math/rand"
	"sync"
	"time"
	"week_5/testcase/client"
)

func main() {
	rand.Seed(time.Now().Unix())
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			client.Client()
			wg.Done()
		}()
	}
	wg.Wait()
}
