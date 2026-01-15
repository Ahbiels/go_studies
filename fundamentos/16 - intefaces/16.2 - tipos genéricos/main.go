package main

import (
	"fmt"
	"strconv"
)

func generico(interf any){ //o any pode substituir o interface{}
	fmt.Println(interf)
}

func generico2(interf ...any){
	fmt.Println(interf...)
}

func parseTypes(input any, target string) any {
	switch result := input.(type){
	case int:
		switch target {
		case "float":
			return float64(result)
		case "string":
			return strconv.Itoa(result)
		default:
			return result
		}
	default:
		return result
	}
}

func main(){
	generico("Olá, mundo!!")
	generico(123)
	generico2(12332, 312,5432,43523,887231,1234536)
	fmt.Printf("%T - %v\n",parseTypes(34, "float"),parseTypes(34, "float"))
	fmt.Printf("%T - %v\n",parseTypes(34, "string"),parseTypes(34, "string"))
}