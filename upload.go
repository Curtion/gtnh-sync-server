package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
)

func upload(w http.ResponseWriter, r *http.Request) {
	var result ResultInfo
	if r.Method == "POST" {
		var req RequestUpload
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err)
		}
		token := req.Token
		if verifyToken(token) {
			log.Println(token, "上传校验成功")
			createFile(token, req.Files)
			result = ResultInfo{Code: 200, Msg: "创建成功", Data: "success"}
		} else {
			log.Println(token, "上传校验失败")
			result = ResultInfo{Code: 403, Msg: "校验失败", Data: "error"}
		}
	} else {
		result = ResultInfo{Code: 400, Msg: "当前仅支持POST方法请求", Data: "error"}
	}
	msg, _ := json.Marshal(result)
	w.Write(msg)
}

func createFile(token string, files []Files) {
	for _, file := range files {
		dir := file.Path
		name := file.Name
		content := file.Content
		fileContent, _ := base64.StdEncoding.DecodeString(content)
		createFolder(token, dir)
		var curPath = path.Join(".", token, dir, name)
		log.Println(curPath, "上传成功")
		f, err := os.OpenFile(curPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
		if err != nil {
			log.Println(err.Error())
		}
		defer f.Close()
		_, err = f.Write(fileContent)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
