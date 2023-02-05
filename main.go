package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	initFolder()
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")
	w.WriteHeader(http.StatusOK)
	result := ResultErr{Code: 401, Msg: "禁止访问", Data: "error"}
	msg, _ := json.Marshal(result)
	w.Write(msg)
}

// Token验证
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

// 初始化文件夹
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

// 创建文件夹
func createFolder(token string, dir string) {
	var curPath = path.Join(".", token, dir)
	if pathExists(curPath) {
		return
	}
	err := os.MkdirAll(curPath, 0777)
	if err != nil {
		log.Println(err)
	}
}

// 判断文件夹是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
