package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func showLog(message string) {
	fmt.Println("-=-=-=-=-=-=-=-=--=")
	fmt.Println(message)
	fmt.Println("-=-=-=-=-=-=-=-=--=")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", listUsers).Methods("GET")
	router.HandleFunc("/usuarios", insertUser).Methods("POST")
	router.HandleFunc("/usuarios/{id}", showUser).Methods("GET")
	router.HandleFunc("/usuarios/{id}", deletUser).Methods("DELETE")
	router.HandleFunc("/usuarios/{id}", updateUser).Methods("PUT")

	fmt.Println("Server on 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
