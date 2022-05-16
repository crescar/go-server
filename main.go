package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type usuario struct {
	username string
	correo   string
	clave    string
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/Login", Login).Methods("POST")
	router.HandleFunc("/Register", Register).Methods("POST")
	router.HandleFunc("/Create", Create).Methods("POST")
	router.HandleFunc("/Tasks", Tasks).Methods("GET")
	router.HandleFunc("/Task/{Id}", Task).Methods("GET")
	router.HandleFunc("/Update/{ID}", UpdateTask).Methods("PUT")
	router.HandleFunc("/Delete/{Id}", DeleteTask).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("peticion recibida")
}
func Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("peticion recibida")
}
func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("peticion recibida")
}
func Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("peticion recibida")
}
func Tasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("peticion recibida")
}
func Task(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("peticion recibida")
}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("peticion recibida")
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("peticion recibida")
}
