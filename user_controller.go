package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var status int
var errors []string

func listUsers(w http.ResponseWriter, r *http.Request) {
	showLog("Listando Usuarios")
	errors = nil
	status = 200

	var users []User

	users = fetchUsers()

	HandleResponse(w, status, users, User{}, errors)
}

func showUser(w http.ResponseWriter, r *http.Request) {
	showLog("Mostrardo Usuario")
	errors = nil
	status = 200

	var user User

	params := mux.Vars(r)
	id := params["id"]

	user = fetchUser(id)

	HandleResponse(w, status, []User{}, user, errors)
}

func deletUser(w http.ResponseWriter, r *http.Request) {
	showLog("Deletando Usuario")
	errors = nil
	status = 200

	var user User

	params := mux.Vars(r)
	id := params["id"]
	db := conn()

	user = fetchUser(id)

	userNotFound := len(errors) > 0

	if !userNotFound {
		_, err := db.Query("DELETE FROM user WHERE id = ?", id)

		if err != nil {
			panic(err)
		}
	}

	HandleResponse(w, status, []User{}, user, errors)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	showLog("Atualizando Usuario")

	errors = nil
	status = 200

	var user User
	var userParams UserRequest

	json.NewDecoder(r.Body).Decode(&userParams)

	nameParams := userParams.Name
	ageParmas := userParams.Age
	emailParams := userParams.Email

	params := mux.Vars(r)
	id := params["id"]
	db := conn()

	validateAttributes(nameParams, ageParmas, emailParams)

	hasError := len(errors) > 0

	if !hasError {
		user = fetchUser(id)

		userNotFound := len(errors) > 0

		if !userNotFound {
			_, err := db.Query("UPDATE user SET name = ?, age = ?, email = ? WHERE id = ?", nameParams, ageParmas, emailParams, id)

			if err != nil {
				panic(err)
			}

			user = fetchUser(id)
		}
	}

	HandleResponse(w, status, []User{}, user, errors)
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	showLog("Inserindo Usuario")
	errors = nil
	status = 200

	var user UserRequest

	json.NewDecoder(r.Body).Decode(&user)

	name := user.Name
	age := user.Age
	email := user.Email

	validateAttributes(name, age, email)

	hasError := len(errors) > 0

	if !hasError {
		db := conn()

		_, err := db.Query("INSERT INTO user (name, age, email) values (?, ?, ?)", name, age, email)

		if err != nil {
			panic(err)
		}
	}

	HandleResponse(w, status, []User{}, User{}, errors)
}

// helper methods

func validateAttributes(name string, age int, email string) {
	if name == "" {
		errors = append(errors, "Nome invalido")
	}
	if age <= 0 {
		errors = append(errors, "Idade invalida")
	}
	if email == "" {
		errors = append(errors, "Email invalido")
	}

	if len(errors) > 0 {
		status = 400
	}
}

func fetchUser(id string) User {
	var userResult User
	db := conn()

	rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		var email string

		err := rows.Scan(&id, &name, &age, &email)

		if err != nil {
			panic(err)
		}

		userResult = User{ID: id, Name: name, Age: age, Email: email}
	}

	if userResult.Name == "" {
		errors = append(errors, "Usuario não encontrado")
		status = 404
	}

	return userResult
}

func fetchUsers() []User {
	var usersResult []User
	db := conn()

	rows, err := db.Query("SELECT * FROM user")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		var email string

		err := rows.Scan(&id, &name, &age, &email)

		if err != nil {
			panic(err)
		}

		usersResult = append(usersResult, User{ID: id, Name: name, Age: age, Email: email})
	}

	if len(usersResult) < 0 {
		errors = append(errors, "Nenhum usuario não encontrado")
		status = 404
	}

	return usersResult
}
