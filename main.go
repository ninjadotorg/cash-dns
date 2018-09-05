package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/ninjadotorg/cash-dns/db"

	"crypto/md5"
	"fmt"
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

	result := map[string]interface{}{}
	result["status"] = 1

	peerUrl := r.URL.Query().Get("peer")
	if peerUrl == "" {
		result["status"] = -1
		result["message"] = "param is invalid"

		json.NewEncoder(w).Encode(result)
		return
	}

	db := db.DB{}
	db.Load("db.json")
	db.Set(fmt.Sprintf("%x", md5.Sum([]byte(peerUrl))), peerUrl)
	db.Save()

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/list", list)
	http.HandleFunc("/register", register)
	log.Println("Server is started on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
