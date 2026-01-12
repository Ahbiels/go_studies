package main

import "fmt"

func main() {
	var age int = 16
	gender := "M"

	if age == 18 {
		fmt.Println("O usuário acabou de se tornar de maior")
	} else if age > 18 {
		fmt.Println("O usuário é maior de idade")
	} else {
		fmt.Println("O usuário é menor de idade")
	}

	fmt.Printf("\n")
	
	if age > 10 {
		fmt.Println("O valor é maior do que 10")
		} else if age == 16 {
			fmt.Println("O valor é 16")
		}
		
	fmt.Printf("\n\n")

	// If init

	if gender_validation := gender; gender_validation == "M" && age >= 18 {
		fmt.Println("Bom dia, senhor!")
	} else {
		fmt.Println("Bom dia, senhora!")
	}


}