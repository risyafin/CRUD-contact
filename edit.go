package main

// import (
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func edit(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Contet-type", "application/json")
// 	if err := r.ParseForm(); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	nama := r.FormValue("nama")
// 	phone := r.FormValue("phone")
	
// 	db, err := connect()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()
// 	_, err = db.Exec("update tb_contact set nama = ?,phone =? where id = ?", nama, phone, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write([]byte("Succes"))

// }
