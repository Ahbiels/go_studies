package main

import (
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	erro := checkmail.ValidateFormat("Angelo@test.dcom")
	erro2 := checkmail.ValidateHost("Angelo@gmail.com")
	fmt.Println(erro)
	fmt.Println(erro2)
}