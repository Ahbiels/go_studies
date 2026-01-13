package main

import (
	"fmt"
)

func soma(values ...int) (total int) {
	for _, value := range values {
		total += value
	}
	return
}

func escrever(text string, values ...int) {
	for _, value := range values {
		fmt.Println(text, value)
	}
}

func main() {
	totalSoma := soma(2,3,5,65,12,12,5,12,6,12,6,1000)
	fmt.Println(totalSoma)
	escrever("Olá mundo", 2,5,23,5,8,7,6,9)
}
