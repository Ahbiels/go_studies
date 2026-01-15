package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) createUser() {
	fmt.Printf("Criando o usuário %v\n", u.name)
}

func (u user) UserMaiorDeIdade() bool{
	return u.age < 18
}

func (u *user) fazerAniversario(){
	u.age++
}

func (u *user) trocarNome(newName string){
	fmt.Printf("\nAlterando o nome do usuário de %v para %v\n", u.name, newName)
	u.name = newName
}

func main() {
	var users []user
	users = append(users, user{"Angelo", 20})
	users = append(users, user{"Caio", 21})
	fmt.Println(users)
	// user1 := user{"Angelo", 20}
	users[0].createUser()
	if users[0].UserMaiorDeIdade() {
		fmt.Println("Usuário é menor de idade")
	} else {
		fmt.Println("Usuário é maior de idade")
	}
	users[0].fazerAniversario()
	users[0].fazerAniversario()

	users[0].trocarNome("João")

	fmt.Println(users[0])
}
