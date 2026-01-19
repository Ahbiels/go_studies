package main

import "fmt"

var max_value int = 45

func main() {
	tasks := make(chan int, max_value)
	results := make(chan int, max_value)

	// vou ter 4 processos que estarão puxando dados dessa fila e fazendo as execuções
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 1; i < max_value; i++ {
		tasks <- i
	}
	close(tasks) //após estiver preenchido, vou fechar o tasks para não receber mais mensagens

	for i := 1; i < max_value; i++ {
		fmt.Println(<-results)
	}
}

// tasks só recebe dados, e o results só envia dados
func worker(tasks chan int, results chan int) {
	for task := range tasks {
		results <- fibonacci(task)
	}
}

func fibonacci(index int) int {
	if index <= 1 {
		return index
	}
	return fibonacci(index-2) + fibonacci(index-1)
}
