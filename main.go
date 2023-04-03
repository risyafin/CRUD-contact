package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type contact struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
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
		w.Write([]byte(err.Error()))
		return
	}
	db, errGorm := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_contact"), &gorm.Config{})
	if errGorm != nil {
		panic("Failed connect to databases")
	}
	result := db.Model(&contact{}).Where("id = ?", id).Updates(map[string]interface{}{
		"nama":     contacts.Nama,
		"phone":    contacts.Phone,
		"username": contacts.Username,
		"password": contacts.Password,
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
	result := db.Select("Nama", "Phone", "Username", "Password").Create(&contact)

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
func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_contact"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var contac contact
	json.NewDecoder(r.Body).Decode(&contac)
	var result contact
	res := db.Where("username = ?", contac.Username).Where("password = ?", contac.Password).First(&result)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("login failed"))
		} else {
			w.Write([]byte(res.Error.Error()))
		}
		return
	}
	claims := MyClaims{
		Username: result.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("Lecang"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println("token :", signedToken)
	w.Write([]byte(signedToken))
}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", HandlerLogin).Methods("POST")
	r.HandleFunc("/users", jwtMiddleware(users)).Methods("GET")
	r.HandleFunc("/users/{id}", jwtMiddleware(user)).Methods("GET")
	r.HandleFunc("/users", jwtMiddleware(create)).Methods("POST")
	r.HandleFunc("/users/{id}", jwtMiddleware(delete)).Methods("DELETE")
	r.HandleFunc("/users/{id}", jwtMiddleware(edit)).Methods("PUT")
	fmt.Println("starting web server at localhost:8080")
	http.ListenAndServe(":8080", r)
}
