package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	fmt.Println(os.Getenv("db_type"))
	fmt.Println(os.Getenv("DB_PASSWORD"))
}