package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	totalOfGoroutines := 10

	fmt.Println("CPUs:", runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(totalOfGoroutines)

	for i := 0; i < totalOfGoroutines; i++ {
		go func() {
			v := count
			runtime.Gosched() //yield do processador entre execuções (trocas)
			v++
			count = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// Sempre que o programa for executado, um valor diferente vai aparecer
// pois cada Goroutine salvou um valor diferente nas variáveis ao mesmo tempo
// go run --race main.go -> verifica quantas condições de corrida há