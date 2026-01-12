package main

import (
	"fmt"
	"time"
)

func forMetodo1(){
	fmt.Println("--------------- Primeiro For ----------------")
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
		time.Sleep(time.Second)
	}
	fmt.Println(i)
}

func forMetodo2(){
	fmt.Println("--------------- Segundo For (mais utilizado) ----------------")
	for x := 0; x<10; x++ {
		fmt.Println(x+1)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n\n")
	for x := 0; x<11; x+=5 {
		fmt.Println(x)
		time.Sleep(time.Second)
	}
}

func forMetodo3(){
	fmt.Println("--------------- For com a clausura range ----------------")
	nomes := [3]string{"Angelo", "Pedro", "João"}
	values := []int{4,2,6,2,7,1,29}
	for indice, value := range nomes[0] {
		fmt.Printf("%v - %v\n", indice, string(value))
	}
	fmt.Printf("\n\n")
	for	value := range 10 {
		if value != 9 {
			fmt.Printf("%v - ",value+1)
		} else {
			fmt.Printf("%v",value+1)
		}
	}
	fmt.Printf("\n\n")
	for _, value := range values {
		fmt.Printf("%v - ",value+1)
	}
}

func forMetodo4(){
	fmt.Println("--------------- For com Map ----------------")
	user := map[string]string{
		"name": "Angelo",
		"surname": "Gabriel",
	}
	for key, value := range user {
		fmt.Printf("%v - %v\n", key, value)
	}
}

func forInfinito() {
	fmt.Println("--------------- For infinito ----------------")
	for {
		fmt.Println("Executando")
		time.Sleep(time.Second)
	}
}

func TestFor() {
	start := time.Now()
	for x := 0; x<10000000; x++ {
		fmt.Println(x)
	}
	exit := time.Since(start)
	fmt.Println(exit)
}

func main() {
	
	// forMetodo1()
	// forMetodo2()
	// forMetodo3()
	// forMetodo4()
	// forInfinito()
	TestFor()
	


}