package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() // -1
	}()
	go func() {
		escrever("Programando em Golang")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() //esperar a contagem das goRountines chegar a 0
}

func escrever(text string) {
	for x := 0; x < 5; x++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
