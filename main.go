package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage endpoint hit")
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("UPDATE")
	myRouter.HandleFunc("/articles", AllPosts).Methods("GET")
	myRouter.HandleFunc("/articles", TestPosts).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	InitDB()
	handleRequests()
	fmt.Println("Go ORM initiated")
}
