package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("É um canal de comunicação que pode receber e enviar dados")
	canal := make(chan string)
	go escrever("Hello, World", canal)

	// for {
	// 	message, open := <-canal //recebendo um valor de dentro do canal
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }

	//ou
	for message2 := range canal { //para cada mensagem que for recebido no canal enquanto ele estiver aberto
		fmt.Println(message2) //mostre na tela
	}

}

func escrever(text string, canal chan string) {
	for x := 0; x < 5; x++ {
		canal <- text //mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}
	close(canal) //evita o deadlock - o canal continua esperando receber dados, sendo que não há mais.
}
