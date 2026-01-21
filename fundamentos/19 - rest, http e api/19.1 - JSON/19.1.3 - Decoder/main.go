package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Users struct {
	Id         int32  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Active     bool   `json:"active"`
	Created_at string `json:"created_at"`
}

func main() {
	userInformations := openFile("file.json")
	var users []Users
	
	if err := json.NewDecoder(strings.NewReader(userInformations)).Decode(&users); err!=nil{
		log.Fatal(err)
	}

	fmt.Println(users)

}

func openFile(nameFile string) string {
	data, err := os.ReadFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
