package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var data []int
	var waitGroup sync.WaitGroup
	// waitGroup.Add(1)
	for x := 0; x < 60; x++ {
		waitGroup.Add(1)
		go func() {
			fmt.Println(x)
			time.Sleep(time.Second)
			data = append(data, x)
			defer waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println(data)
	fmt.Println(len(data))
}
