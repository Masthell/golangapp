package main

import (
	"fmt"
	"net/http"
)

type user struct {
	name string
	age uint16
	money int64
	avggrades float64
	happiness float64
}

func (u user) getAllinfo() string {
	return fmt.Sprintf("User name is: %s. He is %d", u.name, u.age)
}

func (u *user) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := user{"Bob", 25, 58, 2, 3}
	bob.setNewName("alex")
	fmt.Fprintf(w, bob.getAllinfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	// var bob User = ...
	//bob := User{name: "Bob", age: 25, money: -58, avg_grades: 2.3, happiness: 3.2}
	handleRequest()
}

// localhost:8080