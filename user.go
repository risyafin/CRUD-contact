package main

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func user(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("contenct-type", "application/json")
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	db, err := connect()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()
// 	var result = contact{}
// 	err = db.QueryRow("select * from tb_contact where id = (?)", id).Scan(&result.Id, &result.Name, &result.Phone)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	hasil, err := json.Marshal(result)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write(hasil)
// }
