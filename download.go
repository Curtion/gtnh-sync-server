package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func download(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		values := r.URL.Query()
		token := values.Get("token")
		if verifyToken(token) {
			log.Println(token, "下载校验成功")
			result := ResultFiles{Code: 200, Msg: "获取成功", Data: getFiles(token)}
			msg, _ := json.Marshal(result)
			w.Write(msg)
		} else {
			log.Println(token, "下载校验失败")
			result := ResultInfo{Code: 403, Msg: "校验失败", Data: "error"}
			msg, _ := json.Marshal(result)
			w.Write(msg)
		}
	} else {
		result := ResultInfo{Code: 400, Msg: "当前仅支持GET方法请求", Data: "error"}
		msg, _ := json.Marshal(result)
		w.Write(msg)
	}
}

func getFiles(token string) []Files {
	var curPath = path.Join(".", token)
	var files []Files
	filepath.Walk(curPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		file, _ := os.Open(path)
		defer file.Close()
		fileInfo, _ := file.Stat()
		fileSize := fileInfo.Size()
		buffer := make([]byte, fileSize)
		file.Read(buffer)
		fileContent := base64.StdEncoding.EncodeToString(buffer)
		files = append(files, Files{Path: path, Name: info.Name(), Content: fileContent})
		return nil
	})
	return files
}
