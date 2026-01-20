package main

import (
	"fmt"
	"time"
)

func main() {
	ch := escrever("Hello, World")

	for i := 0; i<10; i++{
		fmt.Println(<-ch)
	}
}

func escrever(text string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return ch
}
