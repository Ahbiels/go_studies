package main

import "fmt"

func main(){
	func() {
		fmt.Println("Executando")
	}()

	func(text string){
		fmt.Println(text)
	}("Executando")

	result := func(text string) string {
		return text
	}
	fmt.Println(result("Executando"))
}