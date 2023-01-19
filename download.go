package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
)

func download(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err)
		}
		token := req.Token
		if verifyToken(token) {
			log.Println(token, "验证成功")
			createFolder(req)
		} else {
			log.Println(token, "验证失败")
		}
	} else {
		result := ResultStr{Code: 400, Msg: "当前仅支持POST方法请求", Data: "error"}
		msg, _ := json.Marshal(result)
		w.Write(msg)
	}
}

func createFolder(req Request) {
	var curPath = path.Join(".", req.Token, req.Path)
	if pathExists(curPath) {
		return
	}
	err := os.MkdirAll(curPath, 0777)
	if err != nil {
		log.Println(err)
	}
}

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
