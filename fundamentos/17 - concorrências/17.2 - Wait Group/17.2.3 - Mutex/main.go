package main

// Mutex -> Exclusão mútua
//		-- "Tranca" uma variável ou um trecho de código, fazendo com que somente uma goroutine
// 		-- execute o valor em questão em um dado momento 
//		-- Fiquem na "fila" esperando
//      -- Temos os valores "Lock" e "Unlock"
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

	var mu sync.Mutex

	for i := 0; i < totalOfGoroutines; i++ {
		go func() {
			mu.Lock()
			v := count
			runtime.Gosched() //yield do processador entre execuções (trocas)
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}