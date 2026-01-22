package handlers

import (
	"banco_de_dados/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var table string = "test"

type user struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	FirstDate string `json:"first_date"`
	Ativo     uint8  `json:"ativo"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}
	var create_user user

	if err = json.Unmarshal(body, &create_user); err != nil {
		w.Write([]byte("Erro ao converter usuário para struct"))
		return
	}

	db, err := banco.Conn()

	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()
	// Prepare statement
	statement, err := db.Prepare("INSERT INTO test (name, email, ativo) values (?,?,?)")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(create_user.Name, create_user.Email, create_user.Ativo)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInsert, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Usuário inserido com sucesso! ID: %v", idInsert)

}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conn()

	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM %v", table)
	users_row, err := db.Query(query)

	if err != nil {
		w.Write([]byte("Erro ao buscar usuários no banco"))
		return
	}

	defer users_row.Close()

	var users []user

	for users_row.Next() {
		var user user
		if err := users_row.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.FirstDate,
			&user.Ativo,
		); err != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
		users = append(users, user)
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, err := banco.Conn()

	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	user_row, err := db.Query("SELECT * FROM test WHERE id = ?", id)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	var user user
	if user_row.Next() {
		if err := user_row.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.FirstDate,
			&user.Ativo,
		); err != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var update_user user

	if err = json.Unmarshal(body, &update_user); err != nil {
		w.Write([]byte("Erro ao converter usuário para struct"))
		return
	}

	db, err := banco.Conn()

	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE test SET name = ?, email = ?, ativo = ? WHERE id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(
		update_user.Name,
		update_user.Email,
		update_user.Ativo,
		id,
	); err != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, err := banco.Conn()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM test WHERE id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(
		id,
	); err != nil {
		w.Write([]byte("Erro ao Deleter o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
