package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Hello, World!")
	// go escrever("Programando em Golang") //o executa tudo rapidamente e finaliza
	escrever("Programando em Golang")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
