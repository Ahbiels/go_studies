package maps

import "fmt"

type User struct {
	nome string
	surname string
	age int
	phrase string
}

func Study() {
	fmt.Println("test")

	var users []User

	users = append(users, User{"Angelo", "Gabriel", 20, "Minha casa, minha vida"})

	fmt.Println(users)
}