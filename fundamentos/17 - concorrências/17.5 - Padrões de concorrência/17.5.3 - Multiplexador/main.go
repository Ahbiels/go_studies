package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Hello, World!!"),escrever("Programando em Golang"))
	for i := 1; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <- chan string) <- chan string{
	canalDeSaida := make(chan string)
	go func(){
		for {
			select {
			case mensagem := <- canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <- canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
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
