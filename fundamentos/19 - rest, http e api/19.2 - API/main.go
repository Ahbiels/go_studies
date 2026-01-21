package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
		fmt.Println("Servidor Rodando")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}