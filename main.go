package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type contact struct {
	Id    int    `json:"id"`
	Nama  string `json:"name"`
	Phone string `json:"phone"`
}
func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("contenct-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_contact"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var contacts []contact
	result := db.First(&contacts, id)

	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}

	hasil, err := json.Marshal(contacts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}
func edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var contacts contact
	err := json.NewDecoder(r.Body).Decode(&contacts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	db, errGorm := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_contact"), &gorm.Config{})
	if errGorm != nil {
		panic("Failed connect to databases")
	}
	result := db.Model(&contact{}).Where("id = ?", id).Updates(map[string]interface{}{
		"nama":  contacts.Nama,
		"phone": contacts.Phone,
	})
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("Succes"))
}
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_contact"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	result := db.Delete(&contact{}, id)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("Succes"))
}
func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var contact contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_contact"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	result := db.Select("Nama", "Phone").Create(&contact)

	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("success"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_contact"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var contacts []contact
	result := db.Find(&contacts)

	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	hasil, err := json.Marshal(contacts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", users).Methods("GET")
	r.HandleFunc("/users/{id}", user).Methods("GET")
	r.HandleFunc("/users", create).Methods("POST")
	r.HandleFunc("/users/{id}", delete).Methods("DELETE")
	r.HandleFunc("/users/{id}", edit).Methods("PUT")
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
