package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/ninjadotorg/cash-dns/db"

)

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := db.DB{}
	db.Load("db.json")
	data := db.GetAll()

	listD := []string{}
	for _, v := range data {
		listD = append(listD, v.(string))
	}

	result := map[string]interface{}{}
	result["data"] = listD
	result["status"] = 1

	json.NewEncoder(w).Encode(result)
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	peerUrl := r.URL.Query().Get("peer")
	db := db.DB{}
	db.Load("db.json")
	db.Set(peerUrl, peerUrl)
	db.Save()

	result := map[string]interface{}{}
	result["status"] = 1

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/list", list)
	http.HandleFunc("/register", register)
	log.Println("Server is started on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
