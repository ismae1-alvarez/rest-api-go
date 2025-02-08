package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	port := ":8080"

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/courses", getCourses)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/Users \n")

	io.WriteString(w, "Corrio bien tu endpoint")
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/courses \n")

	io.WriteString(w, "Corrio bien tu endpoint")
}
