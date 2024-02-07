package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go doWork(i, &wg)
	}
	wg.Wait()
	fmt.Println("All complete")
}

func doWork(id int, wg *sync.WaitGroup) {
	i := rand.Intn(5)
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println(id, "Done working after", i, "seconds")
	wg.Done()
}
