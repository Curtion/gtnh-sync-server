package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	initFolder()
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/download", download)
	http.ListenAndServe(":8080", nil)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")
	w.WriteHeader(http.StatusOK)
	result := ResultStr{Code: 403, Msg: "禁止访问", Data: "error"}
	msg, _ := json.Marshal(result)
	w.Write(msg)
}

func verifyToken(token string) bool {
	strPtr, err := os.Open("user.json")
	if err != nil {
		log.Println(err)
	}
	defer strPtr.Close()
	var user User
	json.NewDecoder(strPtr).Decode(&user)
	for _, u := range user {
		if u.Token == token {
			return true
		}
	}
	return false
}

func initFolder() {
	strPtr, err := os.Open("user.json")
	if err != nil {
		log.Println(err)
	}
	defer strPtr.Close()
	var user User
	json.NewDecoder(strPtr).Decode(&user)
	for _, u := range user {
		if pathExists(u.Folder) {
			continue
		}
		err := os.Mkdir(u.Folder, 0777)
		if err != nil {
			log.Println(err)
		}
	}
}
