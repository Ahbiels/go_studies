package main

import (
	"banco_de_dados/banco"
	"banco_de_dados/handlers"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RunServer(port int) {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", handlers.SearchUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.SearchUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods(http.MethodDelete)

	forward_port := ":" + strconv.Itoa(port)

	fmt.Println("Escutando na porta", port)
	log.Fatal(http.ListenAndServe(forward_port, router))
}

func main() {
	RunServer(8023)
	db, err := banco.Conn()

	row, err := db.Query("SELECT * FROM test")

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	// fmt.Println(row)

}
