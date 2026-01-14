package main

import "fmt"

func fibonacci(index uint) uint {
	// fmt.Println(index)
	if index <= 1 {
		return index
	}
	return fibonacci(index-2) + fibonacci(index-1)
}

func exec(value int) int {
	fmt.Println(value)
	if value == 10 {
		return value
	}
	return exec(value + 1)
}

func main() {
	fmt.Println("Hello, World!")
	index := uint(10)
	for i:=uint(0);i<index; i++ {
		fmt.Println(fibonacci(i))
	}
	exec(2)
	// fmt.Println(fibonacci(index))
	// fibonacci(index)
}
