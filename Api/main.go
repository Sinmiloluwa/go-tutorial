package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for users
type User struct {
	UserId  string   `json:"userid"`
	Name    string   `json:"name"`
	Sex     string   `json:"sex"`
	Manager *Manager `json:"manager"`
}

type Manager struct {
	Fullname string `json:"fullname"`
	Sex      string `json:"sex"`
}

var users []User

// middleware
func (c *User) IsEmpty() bool {
	return c.Name == ""
}

func main() {
	fmt.Println("API - Users")
	r := mux.NewRouter()

	// seeding
	users = append(users, User{UserId: "1", Name: "James Emmanuel",
		Sex: "Male", Manager: &Manager{Fullname: "Sinmi Oloyede", Sex: "Male"}})
	users = append(users, User{UserId: "2", Name: "Oluwakemi",
		Sex: "Female", Manager: &Manager{Fullname: "Kehinde Akintola", Sex: "Female"}})

	// routing
	r.HandleFunc("/", server).Methods("GET")
	r.HandleFunc("/users", getAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", getOneUser).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	// port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers

func server(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>My first Api</h1>"))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one user")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, user := range users {
		if user.UserId == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode("No user found")
	return
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one user")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Plase input data")
	}

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.IsEmpty() {
		json.NewEncoder(w).Encode("No data in JSON")
		return
	}

	// generate id, string
	rand.Seed(time.Now().UnixNano())
	user.UserId = strconv.Itoa(rand.Intn(100))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
	return
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one user")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove, add with id
	for index, user := range users {
		if user.UserId == params["id"] {
			users = append(users[:index], users[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.UserId = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a user")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, user := range users {
		if user.UserId == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}

	}
}
